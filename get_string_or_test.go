package req

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetStringOr(t *testing.T) {
	tests := []struct {
		name         string
		req          *http.Request
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "GET value exists",
			req:          httptest.NewRequest("GET", "/path?a=1&b=2", nil),
			key:          "a",
			defaultValue: "default",
			want:         "1",
		},
		{
			name:         "value does not exist returns default",
			req:          httptest.NewRequest("GET", "/path?a=1&b=2", nil),
			key:          "c",
			defaultValue: "default",
			want:         "default",
		},
		{
			name: "POST value exists",
			req: func() *http.Request {
				req := httptest.NewRequest("POST", "/path", nil)
				req.Form = url.Values{"b": {"2"}}
				return req
			}(),
			key:          "b",
			defaultValue: "default",
			want:         "2",
		},
		{
			name: "POST value does not exist",
			req: func() *http.Request {
				req := httptest.NewRequest("POST", "/path", nil)
				req.Form = url.Values{}
				return req
			}(),
			key:          "b",
			defaultValue: "default",
			want:         "default",
		},
		{
			name:         "no value exists",
			req:          httptest.NewRequest("GET", "/path", nil),
			key:          "e",
			defaultValue: "default",
			want:         "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetStringOr(tt.req, tt.key, tt.defaultValue)
			if got != tt.want {
				t.Errorf("GetStringOr() = %q, want %q", got, tt.want)
			}
		})
	}
}
