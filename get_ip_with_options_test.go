package req

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newReq(method, url string, body string) *http.Request {
	r := httptest.NewRequest(method, url, nil)
	return r
}

func TestGetIPWithOptions_PreferForwardedFor(t *testing.T) {
	r := newReq("GET", "/", "")
	r.Header.Set("X-REAL-IP", "203.0.113.10")
	r.Header.Set("X-FORWARDED-FOR", "198.51.100.23, 10.0.0.2") // first is public
	ip := GetIPWithOptions(r, IPOptions{PreferForwardedFor: true, Validate: true, ReturnPrivateIfAllPrivate: true})
	if ip != "198.51.100.23" {
		t.Fatalf("expected 198.51.100.23, got %q", ip)
	}
}

func TestGetIPWithOptions_PreferRealIP(t *testing.T) {
	r := newReq("GET", "/", "")
	r.Header.Set("X-REAL-IP", "203.0.113.10")
	r.Header.Set("X-FORWARDED-FOR", "198.51.100.23, 10.0.0.2")
	ip := GetIPWithOptions(r, IPOptions{PreferForwardedFor: false, Validate: true})
	if ip != "203.0.113.10" {
		t.Fatalf("expected 203.0.113.10, got %q", ip)
	}
}

func TestGetIPWithOptions_TrustedProxies(t *testing.T) {
	// XFF chain: client, proxy1, proxy2 (trusted)
	r := newReq("GET", "/", "")
	r.Header.Set("X-FORWARDED-FOR", "198.51.100.50, 10.0.0.5, 127.0.0.1")
	ip := GetIPWithOptions(r, IPOptions{
		PreferForwardedFor: true,
		TrustedProxies:     []string{"10.0.0.0/8", "127.0.0.1/32"},
		Validate:           true,
	})
	if ip != "198.51.100.50" {
		t.Fatalf("expected 198.51.100.50, got %q", ip)
	}
}

func TestGetIPWithOptions_AllTrusted_ReturnLast(t *testing.T) {
	// All in XFF are trusted; expect last
	r := newReq("GET", "/", "")
	r.Header.Set("X-FORWARDED-FOR", "127.0.0.1, 10.0.0.2")
	ip := GetIPWithOptions(r, IPOptions{
		PreferForwardedFor: true,
		TrustedProxies:     []string{"127.0.0.1/32", "10.0.0.0/8"},
		Validate:           true,
	})
	if ip != "10.0.0.2" {
		t.Fatalf("expected 10.0.0.2, got %q", ip)
	}
}

func TestGetIPWithOptions_PrivateAware_ReturnLastPrivate(t *testing.T) {
	// No trusted proxies; should pick first public if any, else last if ReturnPrivateIfAllPrivate
	r := newReq("GET", "/", "")
	r.Header.Set("X-FORWARDED-FOR", "10.0.0.1, 192.168.1.3, 172.16.0.9")
	ip := GetIPWithOptions(r, IPOptions{
		PreferForwardedFor:        true,
		Validate:                  true,
		ReturnPrivateIfAllPrivate: true,
	})
	if ip != "172.16.0.9" {
		t.Fatalf("expected 172.16.0.9, got %q", ip)
	}
}

func TestGetIPWithOptions_PrivateAware_PickFirstPublic(t *testing.T) {
	r := newReq("GET", "/", "")
	r.Header.Set("X-FORWARDED-FOR", "10.0.0.1, 203.0.113.77, 192.168.1.4")
	ip := GetIPWithOptions(r, IPOptions{
		PreferForwardedFor:        true,
		Validate:                  true,
		ReturnPrivateIfAllPrivate: true,
	})
	if ip != "203.0.113.77" {
		t.Fatalf("expected 203.0.113.77, got %q", ip)
	}
}

func TestGetIPWithOptions_AdditionalHeaders(t *testing.T) {
	r := newReq("GET", "/", "")
	r.Header.Set("CF-Connecting-IP", "203.0.113.200")
	ip := GetIPWithOptions(r, IPOptions{
		AdditionalHeaders: []string{"CF-Connecting-IP"},
		Validate:          true,
	})
	if ip != "203.0.113.200" {
		t.Fatalf("expected 203.0.113.200, got %q", ip)
	}
}

func TestGetIPWithOptions_ValidateSkipsInvalid(t *testing.T) {
	r := newReq("GET", "/", "")
	r.Header.Set("X-FORWARDED-FOR", "not-an-ip, 203.0.113.9")
	ip := GetIPWithOptions(r, IPOptions{
		PreferForwardedFor: true,
		Validate:          true,
	})
	if ip != "203.0.113.9" {
		t.Fatalf("expected 203.0.113.9, got %q", ip)
	}
}

func TestGetIPWithOptions_FallbackRemoteAddr(t *testing.T) {
	r := newReq("GET", "/", "")
	r.RemoteAddr = "203.0.113.5:12345"
	ip := GetIPWithOptions(r, IPOptions{Validate: true})
	if ip != "203.0.113.5" {
		t.Fatalf("expected 203.0.113.5, got %q", ip)
	}
}

func TestGetIPWithOptions_NilRequest(t *testing.T) {
	ip := GetIPWithOptions(nil, IPOptions{})
	if ip != "" {
		t.Fatalf("expected empty string, got %q", ip)
	}
}
