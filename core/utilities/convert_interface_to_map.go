package utilities

import (
	"encoding/json"
	"fmt"
)

func ConvertResultToMap(result interface{}) (map[string]interface{}, error) {
	// Marshal the result to JSON bytes
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error marshalling to JSON: %w", err)
	}

	// Unmarshal the JSON bytes into a map
	var resMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &resMap); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return resMap, nil
}
