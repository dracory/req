package req

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		url      string
		formData url.Values
		headers  map[string]string
		expect   map[string]string
	}{
		{
			name:   "POST with form data",
			method: "POST",
			url:    "http://example.com",
			formData: url.Values{
				"key1": {"value1"},
				"key2": {"value2"},
			},
			headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			expect: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:   "GET with query params",
			method: "GET",
			url:    "http://example.com?key3=value3&key4=value4",
			headers: map[string]string{},
			expect: map[string]string{
				"key3": "value3",
				"key4": "value4",
			},
		},
		{
			name:   "POST with both form and query params",
			method: "POST",
			url:    "http://example.com?key1=query1",
			formData: url.Values{
				"key1": {"form1"}, // Should override query param
				"key2": {"value2"},
			},
			headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			expect: map[string]string{
				"key1": "form1", // Form value should override query param
				"key2": "value2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			req, _ = http.NewRequest(tt.method, tt.url, nil)
			if tt.method == "POST" {
				req, _ = http.NewRequest(tt.method, tt.url, strings.NewReader(tt.formData.Encode()))
			}

			// Set headers
			for k, v := range tt.headers {
				req.Header.Add(k, v)
			}

			result := GetAll(req)

			// Check expected values
			for key, want := range tt.expect {
				if got := result.Get(key); got != want {
					t.Errorf("GetAll() [%s] %s = %v, want %v", tt.name, key, got, want)
				}
			}
		})
	}
}

func TestGetAllGet(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		expect map[string]string
	}{
		{
			name: "Simple query params",
			url:  "http://example.com?key1=value1&key2=value2",
			expect: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Empty query",
			url:  "http://example.com",
			expect: map[string]string{},
		},
		{
			name: "URL-encoded values",
			url:  "http://example.com?name=John%20Doe&email=test%40example.com",
			expect: map[string]string{
				"name":  "John Doe",
				"email": "test@example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", tt.url, nil)
			result := GetAllGet(req)

			// Check expected values exist
			for key, want := range tt.expect {
				if got := result.Get(key); got != want {
					t.Errorf("GetAllGet() [%s] %s = %v, want %v", tt.name, key, got, want)
				}
			}

			// Check no extra values
			if len(result) != len(tt.expect) {
				t.Errorf("GetAllGet() [%s] unexpected number of values: got %d, want %d",
					tt.name, len(result), len(tt.expect))
			}
		})
	}
}

func TestGetAllPost(t *testing.T) {
	tests := []struct {
		name     string
		formData url.Values
		headers  map[string]string
		expect   map[string]string
		hasError bool
	}{
		{
			name: "Simple form data",
			formData: url.Values{
				"username": {"testuser"},
				"password": {"secret"},
			},
			headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			expect: map[string]string{
				"username": "testuser",
				"password": "secret",
			},
			hasError: false,
		},
		{
			name:     "Empty form",
			formData: url.Values{},
			headers: map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			},
			expect:   map[string]string{},
			hasError: false,
		},
		{
			name:     "Missing content type",
			formData: url.Values{"key": {"value"}},
			headers:  map[string]string{},
			expect:   map[string]string{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "http://example.com", strings.NewReader(tt.formData.Encode()))
			
			// Set headers
			for k, v := range tt.headers {
				req.Header.Add(k, v)
			}

			result := GetAllPost(req)

			// Check expected values
			for key, want := range tt.expect {
				if got := result.Get(key); got != want {
					t.Errorf("GetAllPost() [%s] %s = %v, want %v", tt.name, key, got, want)
				}
			}

			// Check no extra values
			if !tt.hasError && len(result) != len(tt.expect) {
				t.Errorf("GetAllPost() [%s] unexpected number of values: got %d, want %d",
					tt.name, len(result), len(tt.expect))
			}
		})
	}
}
