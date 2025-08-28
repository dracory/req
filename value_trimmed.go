package req

import (
	"net/http"
	"strings"
)

// GetStringTrimmed returns a POST or GET key with leading and trailing whitespace removed.
// Returns an empty string if the key does not exist.
//
// Parameters:
//   - r *http.Request: HTTP request
//   - key string: key to get value for
//
// Returns:
//   - string: trimmed value for key, or empty string if not exists
func GetStringTrimmed(r *http.Request, key string) string {
	return strings.TrimSpace(GetString(r, key))
}
