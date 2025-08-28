package req

import "net/http"

// ValueOr returns a POST or GET key, or the provided default value if the key does not exist
// or its value is an empty string.
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to get value for
//  - defaultValue string: default value to return if key does not exist or is empty
//
// Returns:
//  - string: value for key, or default value if key doesn't exist or is empty
func ValueOr(r *http.Request, key string, defaultValue string) string {
	if value := Value(r, key); value != "" {
		return value
	}
	return defaultValue
}
