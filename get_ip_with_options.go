package req

import (
	"net"
	"net/http"
	"strings"
)

// IPOptions configures how GetIPWithOptions determines the client IP.
//
// Behavior overview:
// - Header precedence:
//   When PreferForwardedFor is true, X-Forwarded-For is checked before X-Real-IP.
//   Otherwise, X-Real-IP is checked first (same as GetIP).
// - Trusted proxies:
//   If TrustedProxies is provided (CIDRs or single IPs), the client IP is taken
//   as the first address in X-Forwarded-For that is NOT within the trusted list.
//   If all are trusted, the last address is returned.
// - Additional headers:
//   AdditionalHeaders are checked in order before falling back to RemoteAddr.
// - Validation:
//   If Validate is true, candidate IPs must parse via net.ParseIP; invalid values are skipped.
//
// Note: This function does not mutate request state and does not perform DNS lookups.
// It relies entirely on headers/RemoteAddr and the provided options.
//
// Example trusted proxies list:
//   []string{"127.0.0.1/32", "10.0.0.0/8", "::1/128"}
//
// Common additional headers:
//   []string{"CF-Connecting-IP", "True-Client-IP"}
//
// If you want the simple behavior, use GetIP().
// Use GetIPWithOptions when behind load balancers/reverse-proxies and you control trust.
type IPOptions struct {
	PreferForwardedFor         bool
	TrustedProxies             []string
	AdditionalHeaders          []string
	Validate                   bool
	ReturnPrivateIfAllPrivate  bool // used when no TrustedProxies specified and scanning XFF
}

// GetIPWithOptions determines the client IP using the provided options.
func GetIPWithOptions(r *http.Request, opts IPOptions) string {
	if r == nil {
		return ""
	}

	trustedNets := parseCIDRs(opts.TrustedProxies)

	// Define helpers
	getFirstValid := func(vals ...string) string {
		for _, v := range vals {
			ip := strings.TrimSpace(v)
			if ip == "" {
				continue
			}
			if opts.Validate && net.ParseIP(ip) == nil {
				continue
			}
			return ip
		}
		return ""
	}

	pickFromXFFTrusted := func(xff string) string {
		if xff == "" {
			return ""
		}
		parts := strings.Split(xff, ",")
		last := ""
		for i := 0; i < len(parts); i++ {
			candidate := strings.TrimSpace(parts[i])
			if candidate == "" {
				continue
			}
			if opts.Validate && net.ParseIP(candidate) == nil {
				continue
			}
			last = candidate
			// If not trusted, it's the client IP
			if !containsIP(trustedNets, candidate) {
				return candidate
			}
		}
		return last
	}

	pickFromXFFPrivateAware := func(xff string) string {
		if xff == "" {
			return ""
		}
		parts := strings.Split(xff, ",")
		last := ""
		for _, p := range parts {
			candidate := strings.TrimSpace(p)
			if candidate == "" {
				continue
			}
			if opts.Validate && net.ParseIP(candidate) == nil {
				continue
			}
			last = candidate
			if !IsPrivateIP(candidate) {
				return candidate
			}
		}
		if opts.ReturnPrivateIfAllPrivate {
			return last
		}
		return ""
	}

	getFromXFF := func() string {
		xff := r.Header.Get("X-FORWARDED-FOR")
		if len(trustedNets) > 0 {
			return pickFromXFFTrusted(xff)
		}
		return pickFromXFFPrivateAware(xff)
	}

	getFromRealIP := func() string {
		val := r.Header.Get("X-REAL-IP")
		return getFirstValid(val)
	}

	getFromAdditional := func() string {
		for _, hdr := range opts.AdditionalHeaders {
			if hdr == "" {
				continue
			}
			v := r.Header.Get(hdr)
			if ip := getFirstValid(v); ip != "" {
				return ip
			}
		}
		return ""
	}

	var ip string
	if opts.PreferForwardedFor {
		ip = getFromXFF()
		if ip == "" {
			ip = getFromRealIP()
		}
	} else {
		ip = getFromRealIP()
		if ip == "" {
			ip = getFromXFF()
		}
	}
	if ip == "" {
		ip = getFromAdditional()
	}
	if ip != "" {
		return ip
	}

	// Fallback to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// parseCIDRs parses CIDR strings or single IPs into a slice of *net.IPNet.
func parseCIDRs(vals []string) []*net.IPNet {
	var out []*net.IPNet
	for _, v := range vals {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		if ip := net.ParseIP(v); ip != nil {
			// Convert single IP to /32 or /128
			mask := net.CIDRMask(32, 32)
			if ip.To4() == nil {
				mask = net.CIDRMask(128, 128)
			}
			out = append(out, &net.IPNet{IP: ip, Mask: mask})
			continue
		}
		_, n, err := net.ParseCIDR(v)
		if err == nil && n != nil {
			out = append(out, n)
		}
	}
	return out
}

// containsIP checks whether the given IP string lies within any network in nets.
func containsIP(nets []*net.IPNet, ipStr string) bool {
	ip := net.ParseIP(strings.TrimSpace(ipStr))
	if ip == nil {
		return false
	}
	for _, n := range nets {
		if n.Contains(ip) {
			return true
		}
	}
	return false
}
