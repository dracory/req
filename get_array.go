package req

import (
	"net/http"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

// GetArray retrieves values from the request that match the given key in various formats:
// 1. Direct match (key=value1&key=value2)
// 2. Array notation (key[]=value1&key[]=value2)
// 3. Numbered notation (key[0]=value1&key[1]=value2)
//
// Parameters:
//   - r *http.Request: HTTP request containing the form data
//   - key string: the key to look up in the form data
//   - defaultValue []string: value to return if key is not found
//
// Returns:
//   - []string: array of values for the key, or defaultValue if not found
func GetArray(r *http.Request, key string, defaultValue []string) []string {
	all := GetAll(r)
	if all == nil {
		return defaultValue
	}

	// Check for direct match (key=value1&key=value2)
	if values, exists := all[key]; exists {
		return values
	}

	// Check for array notation (key[]=value1&key[]=value2)
	arrayKey := key + "[]"
	if values, exists := all[arrayKey]; exists {
		return values
	}

	// Check for numbered notation (key[0]=value1&key[1]=value2)
	return extractNumberedValues(all, key)
}

// extractNumberedValues handles the numbered array notation (key[0], key[1], etc.)
func extractNumberedValues(all map[string][]string, key string) []string {
	type indexedValue struct {
		index int
		value string
	}

	var values []indexedValue

	// Find all keys that match the pattern key[number]
	for k, v := range all {
		if !strings.HasPrefix(k, key+"[") || !strings.HasSuffix(k, "]") {
			continue
		}

		// Extract the index from key[index]
		indexStr := strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")
		index := cast.ToInt(indexStr)

		// Get the first value (if any)
		value := ""
		if len(v) > 0 {
			value = v[0]
		}

		values = append(values, indexedValue{index: index, value: value})
	}

	// Sort values by their index
	sort.Slice(values, func(i, j int) bool {
		return values[i].index < values[j].index
	})

	// Extract just the values in order
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.value
	}

	return result
}
