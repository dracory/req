package req

import (
	"net/http"
	"testing"
)

func TestSubdomain(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		expected string
	}{
		{
			name:     "localhost",
			host:     "localhost",
			expected: "",
		},
		{
			name:     "no dot",
			host:     "example",
			expected: "",
		},
		{
			name:     "has dot",
			host:     "sub.example.com",
			expected: "sub",
		},
		{
			name:     "empty host",
			host:     "",
			expected: "",
		},
		{
			name:     "nil host",
			host:     "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "http://"+tt.host, nil)
			if err != nil {
				t.Fatal(err)
			}

			subdomain := Subdomain(req)

			if subdomain != tt.expected {
				t.Errorf("Subdomain() = %q, want %q", subdomain, tt.expected)
			}
		})
	}
}

func TestSubdomain_NilRequest(t *testing.T) {
	subdomain := Subdomain(nil)

	if subdomain != "" {
		t.Errorf("Subdomain() = %q, want \"\"", subdomain)
	}
}
