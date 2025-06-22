package converter

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Convert converts a string value to the specified Go type
func Convert(value string, targetType string) (interface{}, error) {
	// Handle basic types
	switch targetType {
	case "string":
		return value, nil
	case "int":
		return strconv.Atoi(value)
	case "int8":
		val, err := strconv.ParseInt(value, 10, 8)
		return int8(val), err
	case "int16":
		val, err := strconv.ParseInt(value, 10, 16)
		return int16(val), err
	case "int32":
		val, err := strconv.ParseInt(value, 10, 32)
		return int32(val), err
	case "int64":
		return strconv.ParseInt(value, 10, 64)
	case "uint":
		val, err := strconv.ParseUint(value, 10, 64)
		return uint(val), err
	case "uint8":
		val, err := strconv.ParseUint(value, 10, 8)
		return uint8(val), err
	case "uint16":
		val, err := strconv.ParseUint(value, 10, 16)
		return uint16(val), err
	case "uint32":
		val, err := strconv.ParseUint(value, 10, 32)
		return uint32(val), err
	case "uint64":
		return strconv.ParseUint(value, 10, 64)
	case "float32":
		val, err := strconv.ParseFloat(value, 32)
		return float32(val), err
	case "float64":
		return strconv.ParseFloat(value, 64)
	case "bool":
		return strconv.ParseBool(value)
	case "time.Time":
		return parseTime(value)
	case "[]byte":
		return []byte(value), nil
	}

	// Handle slice types
	if strings.HasPrefix(targetType, "[]") {
		return convertSlice(value, targetType[2:])
	}

	// Handle map types
	if strings.HasPrefix(targetType, "map[") {
		return convertMap(value, targetType)
	}

	// Handle pointer types
	if strings.HasPrefix(targetType, "*") {
		return convertPointer(value, targetType[1:])
	}

	// Handle struct types (assume JSON format)
	return convertStruct(value, targetType)
}

// ConvertWithDelimiter converts a string to a slice using a custom delimiter
func ConvertWithDelimiter(value string, targetType string, delimiter string) (interface{}, error) {
	if !strings.HasPrefix(targetType, "[]") {
		return nil, fmt.Errorf("delimiter can only be used with slice types")
	}

	elementType := targetType[2:]
	parts := strings.Split(value, delimiter)

	// Create slice with appropriate type
	switch elementType {
	case "string":
		result := make([]string, len(parts))
		for i, part := range parts {
			result[i] = strings.TrimSpace(part)
		}
		return result, nil
	case "int":
		result := make([]int, len(parts))
		for i, part := range parts {
			val, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return nil, fmt.Errorf("failed to convert element %d: %v", i, err)
			}
			result[i] = val
		}
		return result, nil
	case "float64":
		result := make([]float64, len(parts))
		for i, part := range parts {
			val, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
			if err != nil {
				return nil, fmt.Errorf("failed to convert element %d: %v", i, err)
			}
			result[i] = val
		}
		return result, nil
	case "bool":
		result := make([]bool, len(parts))
		for i, part := range parts {
			val, err := strconv.ParseBool(strings.TrimSpace(part))
			if err != nil {
				return nil, fmt.Errorf("failed to convert element %d: %v", i, err)
			}
			result[i] = val
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported slice element type: %s", elementType)
	}
}

// parseTime attempts to parse time in various formats
func parseTime(value string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
		"2006-01-02",
		time.RFC822,
		time.RFC850,
		time.ANSIC,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, value); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", value)
}

// convertSlice converts a string to a slice type
func convertSlice(value string, elementType string) (interface{}, error) {
	// Try to parse as JSON array first
	var jsonArray []interface{}
	if err := json.Unmarshal([]byte(value), &jsonArray); err == nil {
		return convertJSONSlice(jsonArray, elementType)
	}

	// Fall back to comma-separated values
	return ConvertWithDelimiter(value, "[]"+elementType, ",")
}

// convertJSONSlice converts a JSON array to the target slice type
func convertJSONSlice(jsonArray []interface{}, elementType string) (interface{}, error) {
	switch elementType {
	case "string":
		result := make([]string, len(jsonArray))
		for i, v := range jsonArray {
			if str, ok := v.(string); ok {
				result[i] = str
			} else {
				result[i] = fmt.Sprintf("%v", v)
			}
		}
		return result, nil
	case "int":
		result := make([]int, len(jsonArray))
		for i, v := range jsonArray {
			switch val := v.(type) {
			case float64:
				result[i] = int(val)
			case int:
				result[i] = val
			default:
				return nil, fmt.Errorf("cannot convert %v to int", v)
			}
		}
		return result, nil
	case "float64":
		result := make([]float64, len(jsonArray))
		for i, v := range jsonArray {
			if val, ok := v.(float64); ok {
				result[i] = val
			} else {
				return nil, fmt.Errorf("cannot convert %v to float64", v)
			}
		}
		return result, nil
	case "bool":
		result := make([]bool, len(jsonArray))
		for i, v := range jsonArray {
			if val, ok := v.(bool); ok {
				result[i] = val
			} else {
				return nil, fmt.Errorf("cannot convert %v to bool", v)
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported slice element type: %s", elementType)
	}
}

// convertMap converts a string to a map type
func convertMap(value string, targetType string) (interface{}, error) {
	// Parse as JSON object
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(value), &jsonMap); err != nil {
		return nil, fmt.Errorf("failed to parse map: %v", err)
	}

	// Extract key and value types from targetType (e.g., "map[string]int")
	// This is a simplified implementation - in practice, you'd need more sophisticated parsing
	if strings.Contains(targetType, "string]int") {
		result := make(map[string]int)
		for k, v := range jsonMap {
			if val, ok := v.(float64); ok {
				result[k] = int(val)
			} else {
				return nil, fmt.Errorf("cannot convert map value %v to int", v)
			}
		}
		return result, nil
	}

	if strings.Contains(targetType, "string]string") {
		result := make(map[string]string)
		for k, v := range jsonMap {
			result[k] = fmt.Sprintf("%v", v)
		}
		return result, nil
	}

	return nil, fmt.Errorf("unsupported map type: %s", targetType)
}

// convertPointer converts a string to a pointer type
func convertPointer(value string, targetType string) (interface{}, error) {
	// Convert to the base type first
	result, err := Convert(value, targetType)
	if err != nil {
		return nil, err
	}

	// Create a pointer to the result
	ptr := reflect.New(reflect.TypeOf(result))
	ptr.Elem().Set(reflect.ValueOf(result))
	return ptr.Interface(), nil
}

// convertStruct converts a JSON string to a struct
func convertStruct(value string, targetType string) (interface{}, error) {
	// This is a simplified implementation
	// In practice, you'd need to maintain a registry of struct types
	// or use reflection to create instances dynamically

	// For now, return the JSON as a map
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return nil, fmt.Errorf("failed to parse struct: %v", err)
	}

	return result, nil
}

// GetSupportedTypes returns a list of supported type conversions
func GetSupportedTypes() []string {
	return []string{
		"string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "time.Time", "[]byte",
		"[]string", "[]int", "[]float64", "[]bool",
		"map[string]int", "map[string]string",
		"*int", "*string", "*bool",
	}
}
