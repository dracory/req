package req

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestTrimmedValueOr(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string // Ensure query params here are URL-encoded if needed
		postBody     url.Values
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "GET value exists with leading/trailing spaces",
			method:       "GET",
			url:          "/path?a=%20%20value1%20%20&b=value2",
			key:          "a",
			defaultValue: "default",
			want:         "value1",
		},
		{
			name:         "GET value exists without spaces",
			method:       "GET",
			url:          "/path?a=value1&b=value2",
			key:          "a",
			defaultValue: "default",
			want:         "value1",
		},
		{
			name:         "GET value does not exist, return default",
			method:       "GET",
			url:          "/path?a=1&b=2",
			key:          "c",
			defaultValue: "default_val",
			want:         "default_val",
		},
		{
			name:         "POST value exists with leading/trailing spaces",
			method:       "POST",
			url:          "/path",
			postBody:     url.Values{"b": {"  value2  "}},
			key:          "b",
			defaultValue: "default",
			want:         "value2",
		},
		{
			name:         "POST value exists without spaces",
			method:       "POST",
			url:          "/path",
			postBody:     url.Values{"b": {"value2"}},
			key:          "b",
			defaultValue: "default",
			want:         "value2",
		},
		{
			name:         "POST value does not exist, return default",
			method:       "POST",
			url:          "/path",
			postBody:     url.Values{},
			key:          "b",
			defaultValue: "default_val",
			want:         "default_val",
		},
		{
			name:         "POST value takes precedence over GET and needs trimming",
			method:       "POST",
			url:          "/path?key=get_value",
			postBody:     url.Values{"key": {"  post_value  "}},
			key:          "key",
			defaultValue: "default",
			want:         "post_value",
		},
		{
			name:         "Value contains only spaces, return default",
			method:       "GET",
			url:          "/path?key=%20%20%20",
			key:          "key",
			defaultValue: "default_val",
			want:         "default_val", // TrimSpace("   ") is "", so ValueOr returns default
		},
		{
			name:         "Value is empty string, return default",
			method:       "GET",
			url:          "/path?key=",
			key:          "key",
			defaultValue: "default_val",
			want:         "default_val", // ValueOr returns default
		},
		{
			name:         "No value exists, return default",
			method:       "GET",
			url:          "/path",
			key:          "e",
			defaultValue: "default_val",
			want:         "default_val",
		},
		{
			name:         "Default value has spaces, should be trimmed if returned",
			method:       "GET",
			url:          "/path",
			key:          "e",
			defaultValue: "  spaced_default  ",
			want:         "spaced_default", // The default value itself is trimmed
		},
		{
			name:         "Existing value is only spaces, return trimmed default",
			method:       "GET",
			url:          "/path?key=%20%20",
			key:          "key",
			defaultValue: "  spaced_default  ",
			want:         "spaced_default", // TrimSpace("  ") is "", ValueOr returns default, which is then trimmed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tt.method == "POST" {
				body := ""
				if tt.postBody != nil {
					body = tt.postBody.Encode()
				}
				req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(body))
				if req == nil {
					t.Fatalf("httptest.NewRequest returned nil for POST")
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				err = req.ParseForm()
				if err != nil {
					t.Fatalf("Failed to parse form: %v", err)
				}
			} else { // GET Request
				req = httptest.NewRequest(tt.method, tt.url, nil)
				if req == nil {
					t.Fatalf("httptest.NewRequest returned nil for GET")
				}
				// For GET, ParseForm is implicitly called by FormValue/Value inside ValueOr
			}

			// Call the function under test
			got := TrimmedValueOr(req, tt.key, tt.defaultValue)

			// Assert the result
			if got != tt.want {
				t.Errorf("TrimmedValueOr() = %q, want %q", got, tt.want)
			}
		})
	}
}
