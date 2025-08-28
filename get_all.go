package req

import (
	"maps"
	"net/http"
	"net/url"
)

// GetAll returns all request variables (both GET and POST) as a url.Values object
//
// Parameters:
//   - r *http.Request: HTTP request
//
// Returns:
//   - url.Values: all request variables from both GET and POST
func GetAll(r *http.Request) url.Values {
	gets := GetAllGet(r)
	posts := GetAllPost(r)

	all := url.Values{}

	maps.Copy(all, gets)
	maps.Copy(all, posts)

	return all
}

// GetAllGet returns all GET request variables as a url.Values object
//
// Parameters:
//   - r *http.Request: HTTP request
//
// Returns:
//   - url.Values: GET request variables
func GetAllGet(r *http.Request) url.Values {
	return r.URL.Query()
}

// GetAllPost returns all POST request variables as a url.Values object
// Note: This will only work for application/x-www-form-urlencoded and multipart/form-data
//
// Parameters:
//   - r *http.Request: HTTP request
//
// Returns:
//   - url.Values: POST request variables
func GetAllPost(r *http.Request) url.Values {
	err := r.ParseForm()
	if err != nil {
		return url.Values{}
	}
	return r.PostForm
}

// AllGet returns all GET request variables as a url.Values object.
// Deprecated: prefer GetAllGet for naming consistency. Kept to match documentation.
func AllGet(r *http.Request) url.Values {
	return GetAllGet(r)
}

// AllPost returns all POST request variables as a url.Values object.
// Deprecated: prefer GetAllPost for naming consistency. Kept to match documentation.
func AllPost(r *http.Request) url.Values {
	return GetAllPost(r)
}
