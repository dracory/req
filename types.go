package req

import (
	"net/http"
	"strconv"
)

// GetInt returns the int value of a request parameter.
// Returns 0 if the key is missing or conversion fails.
func GetInt(r *http.Request, key string) int {
	s := GetString(r, key)
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// GetIntOr returns the int value of a request parameter or defaultValue
// if the key is missing or conversion fails.
func GetIntOr(r *http.Request, key string, defaultValue int) int {
	s := GetString(r, key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetInt64 returns the int64 value of a request parameter.
// Returns 0 if the key is missing or conversion fails.
func GetInt64(r *http.Request, key string) int64 {
	s := GetString(r, key)
	if s == "" {
		return 0
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}

// GetInt64Or returns the int64 value of a request parameter or defaultValue
// if the key is missing or conversion fails.
func GetInt64Or(r *http.Request, key string, defaultValue int64) int64 {
	s := GetString(r, key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetFloat64 returns the float64 value of a request parameter.
// Returns 0 if the key is missing or conversion fails.
func GetFloat64(r *http.Request, key string) float64 {
	s := GetString(r, key)
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

// GetFloat64Or returns the float64 value of a request parameter or defaultValue
// if the key is missing or conversion fails.
func GetFloat64Or(r *http.Request, key string, defaultValue float64) float64 {
	s := GetString(r, key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// GetBool returns the bool value of a request parameter.
// Returns false if the key is missing or conversion fails.
func GetBool(r *http.Request, key string) bool {
	s := GetString(r, key)
	if s == "" {
		return false
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}

// GetBoolOr returns the bool value of a request parameter or defaultValue
// if the key is missing or conversion fails.
func GetBoolOr(r *http.Request, key string, defaultValue bool) bool {
	s := GetString(r, key)
	if s == "" {
		return defaultValue
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return v
}
