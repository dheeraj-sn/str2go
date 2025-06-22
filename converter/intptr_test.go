package converter

import (
	"testing"
)

func TestStringToIntPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int
		hasError bool
	}{
		{"positive integer", "123", intPtr(123), false},
		{"negative integer", "-123", intPtr(-123), false},
		{"zero", "0", intPtr(0), false},
		{"large positive", "2147483647", intPtr(2147483647), false},
		{"large negative", "-2147483648", intPtr(-2147483648), false},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
		{"overflow", "9223372036854775808", nil, true},
		{"underflow", "-9223372036854775809", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToIntPtr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToIntPtr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToIntPtr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToIntPtr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToIntPtr(%q) expected non-nil result", tt.input)
				} else {
					intResult := result.(*int)
					if *intResult != *tt.expected {
						t.Errorf("StringToIntPtr(%q) = %v, expected %v", tt.input, *intResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToInt8Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int8
		hasError bool
	}{
		{"positive integer", "123", int8Ptr(123), false},
		{"negative integer", "-123", int8Ptr(-123), false},
		{"zero", "0", int8Ptr(0), false},
		{"max value", "127", int8Ptr(127), false},
		{"min value", "-128", int8Ptr(-128), false},
		{"overflow", "128", nil, true},
		{"underflow", "-129", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt8Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt8Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToInt8Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt8Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToInt8Ptr(%q) expected non-nil result", tt.input)
				} else {
					intResult := result.(*int64)
					if *intResult != int64(*tt.expected) {
						t.Errorf("StringToInt8Ptr(%q) = %v, expected %v", tt.input, *intResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToInt16Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int16
		hasError bool
	}{
		{"positive integer", "123", int16Ptr(123), false},
		{"negative integer", "-123", int16Ptr(-123), false},
		{"zero", "0", int16Ptr(0), false},
		{"max value", "32767", int16Ptr(32767), false},
		{"min value", "-32768", int16Ptr(-32768), false},
		{"overflow", "32768", nil, true},
		{"underflow", "-32769", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt16Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt16Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToInt16Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt16Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToInt16Ptr(%q) expected non-nil result", tt.input)
				} else {
					intResult := result.(*int64)
					if *intResult != int64(*tt.expected) {
						t.Errorf("StringToInt16Ptr(%q) = %v, expected %v", tt.input, *intResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToInt32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int32
		hasError bool
	}{
		{"positive integer", "123", int32Ptr(123), false},
		{"negative integer", "-123", int32Ptr(-123), false},
		{"zero", "0", int32Ptr(0), false},
		{"max value", "2147483647", int32Ptr(2147483647), false},
		{"min value", "-2147483648", int32Ptr(-2147483648), false},
		{"overflow", "2147483648", nil, true},
		{"underflow", "-2147483649", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt32Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt32Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToInt32Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt32Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToInt32Ptr(%q) expected non-nil result", tt.input)
				} else {
					intResult := result.(*int64)
					if *intResult != int64(*tt.expected) {
						t.Errorf("StringToInt32Ptr(%q) = %v, expected %v", tt.input, *intResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToInt64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *int64
		hasError bool
	}{
		{"positive integer", "123", int64Ptr(123), false},
		{"negative integer", "-123", int64Ptr(-123), false},
		{"zero", "0", int64Ptr(0), false},
		{"max value", "9223372036854775807", int64Ptr(9223372036854775807), false},
		{"min value", "-9223372036854775808", int64Ptr(-9223372036854775808), false},
		{"overflow", "9223372036854775808", nil, true},
		{"underflow", "-9223372036854775809", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt64Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt64Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToInt64Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt64Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToInt64Ptr(%q) expected non-nil result", tt.input)
				} else {
					intResult := result.(*int64)
					if *intResult != *tt.expected {
						t.Errorf("StringToInt64Ptr(%q) = %v, expected %v", tt.input, *intResult, *tt.expected)
					}
				}
			}
		})
	}
}

// Helper functions to create int pointers
func intPtr(i int) *int {
	return &i
}

func int8Ptr(i int8) *int8 {
	return &i
}

func int16Ptr(i int16) *int16 {
	return &i
}

func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}
