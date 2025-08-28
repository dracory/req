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
value := req.Value(r, "key")

// Get a value with a default if not found
value := req.ValueOr(r, "key", "default")

// Check if a parameter exists
if req.Has(r, "key") {
    // Parameter exists
}

// Get all parameters from a request
allParams := req.All(r)
```

## Examples

### Getting Values

```go
// Get string value with empty string as default
username := req.Value(r, "username")

// Get value with custom default
age := req.ValueOr(r, "age", "18")

// Get trimmed value (whitespace removed)
searchTerm := req.TrimmedValue(r, "q")
```

### Working with Arrays

```go
// Get array of values
colors := req.Array(r, "colors")

// Check if array contains value
if req.ArrayHas(r, "permissions", "admin") {
    // User has admin permission
}
```

### IP Address Utilities

```go
// Get client IP address
ip := req.IP(r)

// Check if IP is in a private range
if req.IsPrivateIP(ip) {
    // Handle private IP
}
```

### Subdomain Handling

```go
// Extract subdomain from host
subdomain := req.Subdomain("api.example.com") // returns "api"
```

## Available Functions

### Request Parameter Handling
- `Value(r *http.Request, key string) string` - Returns the value of a GET or POST parameter
- `ValueOr(r *http.Request, key string, defaultValue string) string` - Returns a value with a fallback if not found
- `TrimmedValue(r *http.Request, key string) string` - Returns a trimmed (whitespace removed) value
- `TrimmedValueOr(r *http.Request, key string, defaultValue string) string` - Returns a trimmed value with a fallback

### Parameter Existence Checking
- `Has(r *http.Request, key string) bool` - Checks if a parameter exists in GET or POST
- `HasGet(r *http.Request, key string) bool` - Checks if a GET parameter exists
- `HasPost(r *http.Request, key string) bool` - Checks if a POST parameter exists

### Array Operations
- `Array(r *http.Request, key string, defaultValue []string) []string` - Gets an array of values from request parameters
- `ArrayHas(r *http.Request, key string, value string) bool` - Checks if an array contains a specific value

### Map Operations
- `Map(r *http.Request, key string) map[string]string` - Gets a map from request parameters
- `Maps(r *http.Request, key string, defaultValue []map[string]string) []map[string]string` - Gets an array of maps from request parameters

### IP Address Utilities
- `IP(r *http.Request) string` - Gets the client's IP address
- `IsPrivateIP(ip string) bool` - Checks if an IP address is in a private range

### Subdomain Handling
- `Subdomain(host string) string` - Extracts the subdomain from a hostname

### Request Data
- `All(r *http.Request) url.Values` - Gets all request parameters (GET and POST combined)
- `AllGet(r *http.Request) url.Values` - Gets all GET parameters
- `AllPost(r *http.Request) url.Values` - Gets all POST parameters

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.