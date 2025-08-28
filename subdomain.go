package req

import (
	"net/http"
	"strings"
)

// Subdomain finds the subdomain in the host of the given request.
//
// Business logic:
// - extract the host from the request
// - if the host is "localhost", return an empty string
// - if there is no dot in the host, return an empty string
// - otherwise, return the part of the host before the first dot
//
// Parameters:
//   - r (*http.Request): The HTTP request from which to extract the subdomain.
//
// Returns:
//   - string: the subdomain, or an empty string if none found.
//   - error: an error if any, otherwise nil.
func Subdomain(r *http.Request) (string, error) {
	if r == nil || r.URL == nil || r.URL.Host == "" {
		return "", nil
	}

	// If the host is "localhost", there is no subdomain
	if r.URL.Host == "localhost" {
		return "", nil
	}

	// Find the index of the first dot in the host
	i := strings.Index(r.URL.Host, ".")

	// If there is no dot, there is no subdomain
	if i == -1 {
		return "", nil
	}

	// The subdomain is the part of the host before the first dot
	subdomain := r.URL.Host[:i]

	return subdomain, nil
}
