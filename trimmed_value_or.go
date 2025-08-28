package req

import (
	"net/http"
	"strings"
)

// TrimmedValueOr returns a POST or GET key, trimmed, or provided default value if not exists
// or if the found value is empty after trimming. The default value is also trimmed.
//
// Parameters:
//   - r *http.Request: HTTP request
//   - key string: key to get value for
//   - defaultValue string: default value to return if key does not exist or value is empty after trimming
//
// Returns:
//   - string: trimmed value for key, or trimmed default value if appropriate
func TrimmedValueOr(r *http.Request, key string, defaultValue string) string {
	val := Value(r, key)

	trimmedVal := strings.TrimSpace(val)

	if trimmedVal == "" {
		return strings.TrimSpace(defaultValue)
	}

	return trimmedVal
}
