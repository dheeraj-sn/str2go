package converter

import (
	"reflect"
	"testing"
	"time"
)

func TestConvert_BasicTypes(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		targetType string
		expected   interface{}
		expectErr  bool
	}{
		{"string", "hello", "string", "hello", false},
		{"int", "42", "int", 42, false},
		{"int8", "127", "int8", int8(127), false},
		{"int16", "32767", "int16", int16(32767), false},
		{"int32", "2147483647", "int32", int32(2147483647), false},
		{"int64", "9223372036854775807", "int64", int64(9223372036854775807), false},
		{"uint", "42", "uint", uint(42), false},
		{"uint8", "255", "uint8", uint8(255), false},
		{"uint16", "65535", "uint16", uint16(65535), false},
		{"uint32", "4294967295", "uint32", uint32(4294967295), false},
		{"uint64", "18446744073709551615", "uint64", uint64(18446744073709551615), false},
		{"float32", "3.14", "float32", float32(3.14), false},
		{"float64", "3.14159", "float64", 3.14159, false},
		{"bool_true", "true", "bool", true, false},
		{"bool_false", "false", "bool", false, false},
		{"time", "2023-01-02T15:04:05Z", "time.Time", time.Date(2023, 1, 2, 15, 4, 5, 0, time.UTC), false},
		{"bytes", "hello", "[]byte", []byte("hello"), false},
		{"invalid_int", "not_a_number", "int", nil, true},
		{"invalid_float", "not_a_float", "float64", nil, true},
		{"invalid_bool", "not_a_bool", "bool", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.value, tt.targetType)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestConvert_SliceTypes(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		targetType string
		expected   interface{}
		expectErr  bool
	}{
		{"string_slice_json", `["a", "b", "c"]`, "[]string", []string{"a", "b", "c"}, false},
		{"int_slice_json", `[1, 2, 3]`, "[]int", []int{1, 2, 3}, false},
		{"float_slice_json", `[1.1, 2.2, 3.3]`, "[]float64", []float64{1.1, 2.2, 3.3}, false},
		{"bool_slice_json", `[true, false, true]`, "[]bool", []bool{true, false, true}, false},
		{"string_slice_csv", "a,b,c", "[]string", []string{"a", "b", "c"}, false},
		{"int_slice_csv", "1,2,3", "[]int", []int{1, 2, 3}, false},
		{"float_slice_csv", "1.1,2.2,3.3", "[]float64", []float64{1.1, 2.2, 3.3}, false},
		{"bool_slice_csv", "true,false,true", "[]bool", []bool{true, false, true}, false},
		{"invalid_json", "invalid json", "[]string", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.value, tt.targetType)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestConvertWithDelimiter(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		targetType string
		delimiter  string
		expected   interface{}
		expectErr  bool
	}{
		{"string_semicolon", "a;b;c", "[]string", ";", []string{"a", "b", "c"}, false},
		{"int_pipe", "1|2|3", "[]int", "|", []int{1, 2, 3}, false},
		{"float_space", "1.1 2.2 3.3", "[]float64", " ", []float64{1.1, 2.2, 3.3}, false},
		{"bool_tab", "true\tfalse\ttrue", "[]bool", "\t", []bool{true, false, true}, false},
		{"non_slice_type", "42", "int", ",", nil, true},
		{"invalid_element", "1,not_a_number,3", "[]int", ",", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertWithDelimiter(tt.value, tt.targetType, tt.delimiter)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestConvert_MapTypes(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		targetType string
		expected   interface{}
		expectErr  bool
	}{
		{"string_int_map", `{"a": 1, "b": 2}`, "map[string]int", map[string]int{"a": 1, "b": 2}, false},
		{"string_string_map", `{"key1": "value1", "key2": "value2"}`, "map[string]string", map[string]string{"key1": "value1", "key2": "value2"}, false},
		{"invalid_json", "invalid json", "map[string]int", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.value, tt.targetType)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestConvert_PointerTypes(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		targetType   string
		expectedType string
		expectErr    bool
	}{
		{"int_ptr", "42", "*int", "*int", false},
		{"string_ptr", "hello", "*string", "*string", false},
		{"bool_ptr", "true", "*bool", "*bool", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Convert(tt.value, tt.targetType)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			resultType := reflect.TypeOf(result).String()
			if resultType != tt.expectedType {
				t.Errorf("expected type %s, got %s", tt.expectedType, resultType)
			}

			// Check that it's actually a pointer
			if reflect.TypeOf(result).Kind() != reflect.Ptr {
				t.Errorf("expected pointer type, got %s", reflect.TypeOf(result).Kind())
			}
		})
	}
}

func TestConvert_StructTypes(t *testing.T) {
	jsonStr := `{"name": "John", "age": 30, "active": true}`

	result, err := Convert(jsonStr, "Person")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	// Should return a map[string]interface{}
	if _, ok := result.(map[string]interface{}); !ok {
		t.Errorf("expected map[string]interface{}, got %T", result)
	}

	expectedMap := map[string]interface{}{
		"name":   "John",
		"age":    float64(30), // JSON numbers are float64
		"active": true,
	}

	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("expected %v, got %v", expectedMap, result)
	}
}

func TestGetSupportedTypes(t *testing.T) {
	types := GetSupportedTypes()

	expectedTypes := []string{
		"string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "bool", "time.Time", "[]byte",
		"[]string", "[]int", "[]float64", "[]bool",
		"map[string]int", "map[string]string",
		"*int", "*string", "*bool",
	}

	if len(types) != len(expectedTypes) {
		t.Errorf("expected %d types, got %d", len(expectedTypes), len(types))
	}

	for _, expectedType := range expectedTypes {
		found := false
		for _, actualType := range types {
			if actualType == expectedType {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected type %s not found in supported types", expectedType)
		}
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		expected  time.Time
		expectErr bool
	}{
		{"RFC3339", "2023-01-02T15:04:05Z", time.Date(2023, 1, 2, 15, 4, 5, 0, time.UTC), false},
		{"date_only", "2023-01-02", time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"invalid", "invalid_time", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseTime(tt.value)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
