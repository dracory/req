# req

[![Go Reference](https://pkg.go.dev/badge/github.com/dracory/req.svg)](https://pkg.go.dev/github.com/dracory/req)

A Go package providing utility functions for working with HTTP requests, making it easier to handle request parameters, headers, and other common HTTP operations.

## Features

- Retrieve values from GET, POST, and URL parameters with type conversion
- Check for parameter existence in requests
- Work with arrays and maps from request parameters
- IP address utilities
- Subdomain extraction
- Trimmed value handling
- Type-safe value retrieval with fallbacks

## Installation

```bash
go get github.com/dracory/req
```

## Quick Start

```go
import "github.com/dracory/req"

// Get a value from request parameters
value := req.GetString(r, "key")

// Get a value with a default if not found
valueOr := req.GetStringOr(r, "key", "default")

// Check if a parameter exists
if req.Has(r, "key") {
    // Parameter exists
}

// Get all parameters from a request
allParams := req.GetAll(r)
```

## Examples

### Getting Values

```go
// Get string value with empty string as default
username := req.GetString(r, "username")

// Get value with custom default
age := req.GetStringOr(r, "age", "18")

// Get trimmed value (whitespace removed)
searchTerm := req.GetStringTrimmed(r, "q")
```

### Working with Arrays

```go
// Get array of values
colors := req.GetArray(r, "colors", nil)

// Check if array contains a value (simple contains check)
hasAdmin := false
for _, v := range colors {
    if v == "admin" { hasAdmin = true; break }
}
if hasAdmin {
    // User has admin permission
}

```

### IP Address Utilities

```go
// Get client IP address
ip := req.GetIP(r)

// Check if IP is in a private range
if req.IsPrivateIP(ip) {
    // Handle private IP
}
```

### Subdomain Handling

```go
// Extract subdomain from host
subdomain := req.GetSubdomain(r) // returns "api" for host like api.example.com
```

## Available Functions

### Request Parameter Handling
- `GetString(r *http.Request, key string) string` - Returns the value of a GET or POST parameter
- `GetStringOr(r *http.Request, key string, defaultValue string) string` - Returns a value with a fallback if not found
- `GetStringTrimmed(r *http.Request, key string) string` - Returns a trimmed (whitespace removed) value
- `GetStringTrimmedOr(r *http.Request, key string, defaultValue string) string` - Returns a trimmed value with a fallback

### Parameter Existence Checking
- `Has(r *http.Request, key string) bool` - Checks if a parameter exists in GET or POST
- `HasGet(r *http.Request, key string) bool` - Checks if a GET parameter exists
- `HasPost(r *http.Request, key string) bool` - Checks if a POST parameter exists

### Array Operations
- `GetArray(r *http.Request, key string, defaultValue []string) []string` - Gets an array of values from request parameters

### Map Operations
- `GetMap(r *http.Request, key string) map[string]string` - Gets a map from request parameters
- `GetMaps(r *http.Request, key string, defaultValue []map[string]string) []map[string]string` - Gets an array of maps from request parameters

### IP Address Utilities
- `GetIP(r *http.Request) string` - Gets the client's IP address
- `IsPrivateIP(ip string) bool` - Checks if an IP address is in a private range

### Subdomain Handling
- `GetSubdomain(r *http.Request) string` - Extracts the subdomain from the request hostname

### Request Data
- `GetAll(r *http.Request) url.Values` - Gets all request parameters (GET and POST combined)
- `GetAllGet(r *http.Request) url.Values` - Gets all GET parameters
- `GetAllPost(r *http.Request) url.Values` - Gets all POST parameters

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.