package req_test // Use _test package convention

import (
	"net/http/httptest"
	"testing"

	"github.com/dracory/base/req" // Adjust import path if necessary
)

func TestIP(t *testing.T) {
	tests := []struct {
		name          string
		headers       map[string]string
		remoteAddr    string
		expectedIP    string
		expectError   bool // Although the current IP func doesn't return error, good practice
		errorContains string
	}{
		{
			name: "X-REAL-IP present",
			headers: map[string]string{
				"X-REAL-IP": "192.168.1.1",
			},
			remoteAddr: "10.0.0.1:12345",
			expectedIP: "192.168.1.1",
		},
		{
			name: "X-FORWARDED-FOR present (single IP)",
			headers: map[string]string{
				"X-FORWARDED-FOR": "203.0.113.1",
			},
			remoteAddr: "10.0.0.1:12345",
			expectedIP: "203.0.113.1",
		},
		{
			name: "X-FORWARDED-FOR present (multiple IPs)",
			headers: map[string]string{
				"X-FORWARDED-FOR": "203.0.113.5, 192.168.1.100, 10.0.0.5", // Should take the first one
			},
			remoteAddr: "10.0.0.1:12345",
			expectedIP: "203.0.113.5",
		},
		{
			name: "X-FORWARDED-FOR present (multiple IPs with spaces)",
			headers: map[string]string{
				"X-FORWARDED-FOR": " 203.0.113.6 , 192.168.1.101 ", // Should take the first one after trimming
			},
			remoteAddr: "10.0.0.1:12345",
			expectedIP: " 203.0.113.6 ", // Current implementation returns the first element as is, including spaces
		},
		{
			name:       "Only RemoteAddr present (IPv4)",
			headers:    map[string]string{},
			remoteAddr: "172.16.0.1:54321",
			expectedIP: "172.16.0.1",
		},
		{
			name:       "Only RemoteAddr present (IPv6)",
			headers:    map[string]string{},
			remoteAddr: "[2001:db8::1]:8080",
			expectedIP: "2001:db8::1",
		},
		{
			name:        "Only RemoteAddr present (No Port)",
			headers:     map[string]string{},
			remoteAddr:  "192.0.2.1", // net.SplitHostPort expects host:port
			expectedIP:  "",          // Will fail SplitHostPort
			expectError: true,        // Implicitly, as SplitHostPort fails
		},
		{
			name:        "Invalid RemoteAddr",
			headers:     map[string]string{},
			remoteAddr:  "invalid-address",
			expectedIP:  "",
			expectError: true, // Implicitly, as SplitHostPort fails
		},
		{
			name:        "No headers, empty RemoteAddr",
			headers:     map[string]string{},
			remoteAddr:  "",
			expectedIP:  "",
			expectError: true, // Implicitly, as SplitHostPort fails
		},
		{
			name: "X-REAL-IP takes precedence over X-FORWARDED-FOR",
			headers: map[string]string{
				"X-REAL-IP":       "192.168.10.10",
				"X-FORWARDED-FOR": "203.0.113.10",
			},
			remoteAddr: "10.0.0.1:12345",
			expectedIP: "192.168.10.10",
		},
		{
			name: "X-FORWARDED-FOR takes precedence over RemoteAddr",
			headers: map[string]string{
				"X-FORWARDED-FOR": "203.0.113.20",
			},
			remoteAddr: "172.16.10.10:54321",
			expectedIP: "203.0.113.20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock request
			r := httptest.NewRequest("GET", "http://example.com", nil)

			// Set headers
			for key, value := range tt.headers {
				r.Header.Set(key, value)
			}

			// Set RemoteAddr
			r.RemoteAddr = tt.remoteAddr

			// Call the function under test
			gotIP := req.IP(r) // Assuming IP is a method on req package, not http.Request

			// Assertions
			if gotIP != tt.expectedIP {
				t.Errorf("IP() got = %q, want %q", gotIP, tt.expectedIP)
			}

			// Note: The current IP function doesn't return an error,
			// so error checking isn't directly applicable unless the function signature changes.
			// The `expectError` flag is more for documenting cases where an internal error occurs (like SplitHostPort failing).
		})
	}
}

// Test case for nil request, although the current implementation doesn't handle it explicitly
// and would panic. A robust implementation might check for nil.
func TestIP_NilRequest(t *testing.T) {
	// It's generally better practice for functions accepting pointers to check for nil.
	// We expect a panic here based on the current code.
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("IP(nil) did not panic, but was expected to")
		} else {
			t.Logf("Caught expected panic: %v", r)
		}
	}()

	// This line will cause a panic because the function tries to access req.Header
	_ = req.IP(nil)
}
