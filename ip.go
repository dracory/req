package req

import (
	"net"
	"net/http"
	"strings"
)

// IsPrivateIP checks if an IP address is in a private network range.
// It supports both IPv4 and IPv6 addresses.
//
// Parameters:
//   - ipStr: The IP address to check (can be in string format)
//
// Returns:
//   - bool: true if the IP is in a private range, false otherwise
func IsPrivateIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	// Check for IPv4 private ranges
	if ip4 := ip.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10: // 10.0.0.0/8
			return true
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31: // 172.16.0.0/12
			return true
		case ip4[0] == 192 && ip4[1] == 168: // 192.168.0.0/16
			return true
		case ip4[0] == 100 && (ip4[1]&0b11000000) == 64: // 100.64.0.0/10 (Carrier-grade NAT)
			return true
		default:
			return false
		}
	}

	// Check for IPv6 private ranges (ULA - Unique Local Addresses)
	if ip.To16() != nil {
		// fc00::/7 - Unique Local Addresses (includes fc00::/8 and fd00::/8)
		if len(ip) == net.IPv6len && ip[0]&0xfe == 0xfc {
			return true
		}
		// fe80::/10 - Link-local addresses
		if len(ip) == net.IPv6len && ip[0] == 0xfe && (ip[1]&0xc0) == 0x80 {
			return true
		}
	}

	return false
}

// IP gets the IP address for the user by checking X-REAL-IP, X-FORWARDED-FOR headers,
// and finally falling back to RemoteAddr. For X-FORWARDED-FOR, it returns the first
// non-private IP address in the chain, or the last IP if all are private.
func IP(r *http.Request) string {
	// Get IP from the X-REAL-IP header
	realIP := r.Header.Get("X-REAL-IP")
	if realIP != "" {
		return realIP
	}

	// Get IP from X-FORWARDED-FOR header
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		// Split the header value by commas and trim whitespace
		splitIps := strings.Split(forwarded, ",")
		var lastIP string

		// Iterate through all IPs in the X-FORWARDED-FOR header
		for _, ipStr := range splitIps {
			ipStr = strings.TrimSpace(ipStr)
			if ipStr == "" {
				continue
			}

			// Store the last valid IP in case all are private
			lastIP = ipStr

			// Skip private IPs
			if IsPrivateIP(ipStr) {
				continue
			}

			// Return the first non-private IP
			return ipStr
		}

		// If we get here, all IPs were private or empty, return the last one
		if lastIP != "" {
			return lastIP
		}
	}

	// Fall back to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		// If SplitHostPort fails, try using RemoteAddr as is
		return r.RemoteAddr
	}

	return host
}
