package req

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestValueTrimmed(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		url      string // Ensure query params here are URL-encoded if needed
		postBody url.Values
		key      string
		want     string
		setupReq func() *http.Request
	}{
		{
			name:   "GET value exists with leading/trailing spaces",
			method: "GET",
			// Corrected: Spaces in query parameters must be URL encoded
			url:  "/path?a=%20%20value1%20%20&b=value2",
			key:  "a",
			want: "value1",
		},
		{
			name:   "GET value exists without spaces",
			method: "GET",
			url:    "/path?a=value1&b=value2",
			key:    "a",
			want:   "value1",
		},
		{
			name:   "GET value does not exist",
			method: "GET",
			url:    "/path?a=1&b=2",
			key:    "c",
			want:   "",
		},
		{
			name:     "POST value exists with leading/trailing spaces",
			method:   "POST",
			url:      "/path",
			postBody: url.Values{"b": {"  value2  "}}, // url.Values.Encode() handles encoding
			key:      "b",
			want:     "value2",
		},
		{
			name:     "POST value exists without spaces",
			method:   "POST",
			url:      "/path",
			postBody: url.Values{"b": {"value2"}},
			key:      "b",
			want:     "value2",
		},
		{
			name:     "POST value does not exist",
			method:   "POST",
			url:      "/path",
			postBody: url.Values{},
			key:      "b",
			want:     "",
		},
		{
			name:     "POST value takes precedence over GET and needs trimming",
			method:   "POST",
			url:      "/path?key=get_value", // GET part is fine
			postBody: url.Values{"key": {"  post_value  "}},
			key:      "key",
			want:     "post_value",
		},
		{
			name:   "Value contains only spaces",
			method: "GET",
			// Corrected: Spaces in query parameters must be URL encoded
			url:  "/path?key=%20%20%20",
			key:  "key",
			want: "",
		},
		{
			name:   "Value is empty string",
			method: "GET",
			url:    "/path?key=",
			key:    "key",
			want:   "",
		},
		{
			name:   "No value exists",
			method: "GET",
			url:    "/path",
			key:    "e",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error // Declare err variable

			if tt.setupReq != nil {
				req = tt.setupReq()
			} else if tt.method == "POST" {
				// Create POST request
				body := ""
				if tt.postBody != nil {
					body = tt.postBody.Encode() // url.Values.Encode handles spacing correctly
				}
				req = httptest.NewRequest(tt.method, tt.url, strings.NewReader(body))
				if req == nil {
					t.Fatalf("httptest.NewRequest returned nil for POST") // Added nil check
				}
				// Need to set Content-Type for ParseForm to work correctly for POST
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				// Manually parse form for POST tests, as Value calls r.FormValue which calls ParseForm/ParseMultipartForm
				err = req.ParseForm() // Assign to err
				if err != nil {
					t.Fatalf("Failed to parse form: %v", err)
				}
			} else { // GET Request
				// Create GET request - tt.url should now have correctly encoded params
				req = httptest.NewRequest(tt.method, tt.url, nil)
				if req == nil {
					t.Fatalf("httptest.NewRequest returned nil for GET") // Added nil check
				}
				// No need to manually call ParseForm for GET; r.FormValue handles it.
			}

			// Call the function under test
			got := ValueTrimmed(req, tt.key)

			// Assert the result
			if got != tt.want {
				t.Errorf("ValueTrimmed() = %q, want %q", got, tt.want)
			}
		})
	}
}
