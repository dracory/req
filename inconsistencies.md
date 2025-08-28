# Inconsistencies and Areas for Improvement

This document outlines inconsistencies, bugs, and areas for improvement found in the `req` package.

## 1. Inconsistent `Has` functions

The `HasPost` function does not behave as its name suggests. It checks for the existence of a key in `r.Form`, which is populated by `r.ParseForm()`. `r.ParseForm()` parses both the request body and the query string, so `HasPost` ends up checking for both POST and GET parameters.

**Recommendation:**

To correctly check for only POST parameters, `HasPost` should use `r.PostForm`.

```go
// HasPost returns true if POST key exists
func HasPost(r *http.Request, key string) bool {
    r.ParseForm() // Or r.ParseMultipartForm if needed
    _, exists := r.PostForm[key]
    return exists
}
```

## 2. `Subdomain` function signature

The `Subdomain` function is declared to return an `(string, error)`, but it never returns a non-nil error.

**Recommendation:**

The function signature should be changed to `func Subdomain(r *http.Request) string` to be consistent with the other functions in the package and to reflect its actual behavior.

## 3. `Array` function complexity

The `Array` function is very complex, especially the sorting logic. It can be made more readable and maintainable.

**Recommendation:**

Refactor the `Array` function to simplify the logic for handling different array formats and for sorting the values.

## 4. Bugs in `maps.go`

### 4.1. `filterKeyEntries` function

There is a bug in the `filterKeyEntries` function. The line `key, _ := split[0], split[1]` uses `key` which shadows the function's `key` parameter. Also, the second part of the split, `split[1]`, is not being used. This suggests the logic for parsing map arrays is incorrect.

### 4.2. `Maps` function

The `Maps` function assumes that all value slices in the `keyEntries` map have the same length (`lenValues := len(keyEntries[keys[0]])`). If the lengths are different, this will cause a panic.

**Recommendation:**

The logic in `maps.go` needs to be reviewed and fixed to correctly and safely parse arrays of maps from the request.

## 5. Redundancy in `Value` and `ValueOr`

The code for retrieving a value from the request is duplicated in `Value` and `ValueOr`.

**Recommendation:**

`ValueOr` should call `Value` and then return the default value if the result is empty.

```go
func ValueOr(r *http.Request, key string, defaultValue string) string {
    value := Value(r, key)
    if value == "" {
        return defaultValue
    }
    return value
}
```

## 6. Incorrect `IP` logic

The `IP` function, when parsing the `X-FORWARDED-FOR` header, splits the string by commas but then immediately returns the first IP address in the list. The purpose of iterating is defeated.

**Recommendation:**

The logic should be updated to iterate through the IP addresses and return the first one that is a valid, non-private IP.

## 7. Missing functions

The `doc.go` file mentions `AllGet(r)` and `AllPost(r)` functions, but these are not implemented in the package.

**Recommendation:**

Implement the `AllGet` and `AllPost` functions to match the documentation.

```go
// AllGet returns all GET request variables as a url.Values object
func AllGet(r *http.Request) url.Values {
    return r.URL.Query()
}

// AllPost returns all POST request variables as a url.Values object
func AllPost(r *http.Request) url.Values {
    r.ParseForm()
    return r.PostForm
}
```

## 8. Unexpected behavior in `TrimmedValueOr`

The `TrimmedValueOr` function trims the provided `defaultValue`. This might be unexpected for a user of the function.

**Recommendation:**

Document this behavior clearly in the function's comment. The current comment is good, but it could be made more explicit that the default value is also trimmed.
