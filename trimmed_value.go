package req

import (
	"net/http"
	"strings"
)

// TrimmedValue returns a POST or GET key, trimmed, or empty string if not exists
//
// Parameters:
//   - r *http.Request: HTTP request
//   - key string: key to get value for
//
// Returns:
//   - string: trimmed value for key, or empty string if not exists
func TrimmedValue(r *http.Request, key string) string {
	return strings.TrimSpace(Value(r, key))
}
