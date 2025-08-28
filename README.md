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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.