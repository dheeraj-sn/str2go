package converter

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasError bool
	}{
		{"positive integer", "123", 123, false},
		{"negative integer", "-123", -123, false},
		{"zero", "0", 0, false},
		{"large positive", "2147483647", 2147483647, false},
		{"large negative", "-2147483648", -2147483648, false},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
		{"overflow", "9223372036854775808", 0, true},
		{"underflow", "-9223372036854775809", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt(%q) unexpected error: %v", tt.input, err)
				}
				intResult := result.(int)
				if intResult != tt.expected {
					t.Errorf("StringToInt(%q) = %v, expected %v", tt.input, intResult, tt.expected)
				}
			}
		})
	}
}

func TestStringToInt8(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int8
		hasError bool
	}{
		{"positive integer", "123", 123, false},
		{"negative integer", "-123", -123, false},
		{"zero", "0", 0, false},
		{"max value", "127", 127, false},
		{"min value", "-128", -128, false},
		{"overflow", "128", 0, true},
		{"underflow", "-129", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt8(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt8(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt8(%q) unexpected error: %v", tt.input, err)
				}
				intResult := result.(int64)
				if intResult != int64(tt.expected) {
					t.Errorf("StringToInt8(%q) = %v, expected %v", tt.input, intResult, tt.expected)
				}
			}
		})
	}
}

func TestStringToInt16(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int16
		hasError bool
	}{
		{"positive integer", "123", 123, false},
		{"negative integer", "-123", -123, false},
		{"zero", "0", 0, false},
		{"max value", "32767", 32767, false},
		{"min value", "-32768", -32768, false},
		{"overflow", "32768", 0, true},
		{"underflow", "-32769", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt16(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt16(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt16(%q) unexpected error: %v", tt.input, err)
				}
				intResult := result.(int64)
				if intResult != int64(tt.expected) {
					t.Errorf("StringToInt16(%q) = %v, expected %v", tt.input, intResult, tt.expected)
				}
			}
		})
	}
}

func TestStringToInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int32
		hasError bool
	}{
		{"positive integer", "123", 123, false},
		{"negative integer", "-123", -123, false},
		{"zero", "0", 0, false},
		{"max value", "2147483647", 2147483647, false},
		{"min value", "-2147483648", -2147483648, false},
		{"overflow", "2147483648", 0, true},
		{"underflow", "-2147483649", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt32(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt32(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt32(%q) unexpected error: %v", tt.input, err)
				}
				intResult := result.(int64)
				if intResult != int64(tt.expected) {
					t.Errorf("StringToInt32(%q) = %v, expected %v", tt.input, intResult, tt.expected)
				}
			}
		})
	}
}

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		hasError bool
	}{
		{"positive integer", "123", 123, false},
		{"negative integer", "-123", -123, false},
		{"zero", "0", 0, false},
		{"max value", "9223372036854775807", 9223372036854775807, false},
		{"min value", "-9223372036854775808", -9223372036854775808, false},
		{"overflow", "9223372036854775808", 0, true},
		{"underflow", "-9223372036854775809", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt64(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToInt64(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToInt64(%q) unexpected error: %v", tt.input, err)
				}
				intResult := result.(int64)
				if intResult != tt.expected {
					t.Errorf("StringToInt64(%q) = %v, expected %v", tt.input, intResult, tt.expected)
				}
			}
		})
	}
}
