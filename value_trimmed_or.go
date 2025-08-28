package req

import (
	"net/http"
	"strings"
)

// GetStringTrimmedOr returns a POST or GET key with leading and trailing whitespace removed.
// If the resulting value is empty or the key doesn't exist, returns the provided default value.
// Note: The default value is also trimmed of leading and trailing whitespace.
//
// Parameters:
//   - r *http.Request: HTTP request
//   - key string: key to get value for
//   - defaultValue string: default value to return if key doesn't exist or value is empty after trimming
//
// Returns:
//   - string: trimmed value for key, or trimmed default value if the key doesn't exist or value is empty after trimming
func GetStringTrimmedOr(r *http.Request, key string, defaultValue string) string {
	val := GetString(r, key)
	trimmedVal := strings.TrimSpace(val)

	if trimmedVal == "" {
		return strings.TrimSpace(defaultValue)
	}

	return trimmedVal
}
