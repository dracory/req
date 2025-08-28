package req

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestValue(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		key  string
		want string
	}{
		{
			name: "GET value exists",
			req:  httptest.NewRequest("GET", "/path?a=1&b=2", nil),
			key:  "a",
			want: "1",
		},
		{
			name: "GET value does not exist",
			req:  httptest.NewRequest("GET", "/path?a=1&b=2", nil),
			key:  "c",
			want: "",
		},
		{
			name: "POST value exists",
			req: func() *http.Request {
				req := httptest.NewRequest("POST", "/path", nil)
				req.Form = url.Values{"b": {"2"}}
				return req
			}(),
			key:  "b",
			want: "2",
		},
		{
			name: "POST value does not exist",
			req: func() *http.Request {
				req := httptest.NewRequest("POST", "/path", nil)
				req.Form = url.Values{}
				return req
			}(),
			key:  "b",
			want: "",
		},
		{
			name: "no value exists",
			req:  httptest.NewRequest("GET", "/path", nil),
			key:  "e",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Value(tt.req, tt.key)
			if got != tt.want {
				t.Errorf("Value() = %q, want %q", got, tt.want)
			}
		})
	}
}
