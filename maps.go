package req

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Maps parses an array of maps from request parameters with the given key prefix.
// The expected format is key[mapKey1][mapKey2] = value.
//
// Parameters:
//   - r: The HTTP request
//   - key: The base key to look for in the request
//   - defaultValue: The default value to return if no matching parameters are found
//
// Returns:
//   - []map[string]string: An array of maps containing the parsed values
func Maps(r *http.Request, key string, defaultValue []map[string]string) []map[string]string {
	all := GetAll(r)

	if all == nil {
		return defaultValue
	}

	// Get all entries with the given key prefix
	keyEntries, err := filterKeyEntries(all, key)
	if err != nil {
		return defaultValue
	}

	if len(keyEntries) == 0 {
		return defaultValue
	}

	// Find the maximum number of values for any key
	var maxValues int
	for _, values := range keyEntries {
		if len(values) > maxValues {
			maxValues = len(values)
		}
	}

	// If no values found, return default
	if maxValues == 0 {
		return defaultValue
	}

	// Build the result array
	result := make([]map[string]string, 0, maxValues)

	for i := 0; i < maxValues; i++ {
		m := make(map[string]string, len(keyEntries))
		for k, values := range keyEntries {
			// If we have fewer values for this key, use an empty string
			if i < len(values) {
				m[k] = values[i]
			} else {
				m[k] = ""
			}
		}
		result = append(result, m)
	}

	return result
}

// filterKeyEntries extracts map entries from url.Values that match the given key pattern.
// The expected format is key[mapKey1][mapKey2] = value.
// Returns a map where the key is the map key and the value is the array of values for that key.
// Returns an error if no valid entries are found.
func filterKeyEntries(all url.Values, key string) (map[string][]string, error) {
	result := make(map[string][]string)
	prefix := key + "["
	hasEntries := false

	for k, values := range all {
		// Skip if not in the expected format
		if !strings.HasPrefix(k, prefix) || !strings.HasSuffix(k, "]") {
			continue
		}

		// Extract the part between key[ and ]
		suffix := k[len(prefix) : len(k)-1]
		if suffix == "" {
			continue
		}

		// Split into map keys
		mapKeys := strings.Split(suffix, "][")
		if len(mapKeys) != 2 || mapKeys[0] == "" || mapKeys[1] == "" {
			continue
		}

		// Use the second key as the map key
		mapKey := mapKeys[1]
		result[mapKey] = values
		hasEntries = true
	}

	if !hasEntries {
		return nil, fmt.Errorf("no valid map entries found for key %s", key)
	}

	return result, nil
}
