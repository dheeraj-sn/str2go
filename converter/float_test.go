package converter

import (
	"math"
	"testing"
)

func TestStringToFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float32
		hasError bool
	}{
		{"positive integer", "123", 123.0, false},
		{"negative integer", "-123", -123.0, false},
		{"positive decimal", "123.456", 123.456, false},
		{"negative decimal", "-123.456", -123.456, false},
		{"zero", "0", 0.0, false},
		{"zero decimal", "0.0", 0.0, false},
		{"scientific notation", "1.23e2", 123.0, false},
		{"negative scientific notation", "-1.23e2", -123.0, false},
		{"large number", "999999.999", 999999.999, false},
		{"invalid string", "invalid", 0.0, true},
		{"empty string", "", 0.0, true},
		{"overflow", "1e100", 0.0, true},
		{"underflow", "-1e100", 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToFloat32(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToFloat32(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToFloat32(%q) unexpected error: %v", tt.input, err)
				}
				// Type assert the result to float64 (since ParseFloat returns float64)
				floatResult := result.(float64)
				// Use approximate comparison for floating point values
				if math.Abs(floatResult-float64(tt.expected)) > 1e-6 {
					t.Errorf("StringToFloat32(%q) = %v, expected %v", tt.input, floatResult, tt.expected)
				}
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		hasError bool
	}{
		{"positive integer", "123", 123.0, false},
		{"negative integer", "-123", -123.0, false},
		{"positive decimal", "123.456", 123.456, false},
		{"negative decimal", "-123.456", -123.456, false},
		{"zero", "0", 0.0, false},
		{"zero decimal", "0.0", 0.0, false},
		{"scientific notation", "1.23e2", 123.0, false},
		{"negative scientific notation", "-1.23e2", -123.0, false},
		{"large number", "999999.999", 999999.999, false},
		{"very large number", "1e308", 1e308, false},
		{"very small number", "1e-308", 1e-308, false},
		{"invalid string", "invalid", 0.0, true},
		{"empty string", "", 0.0, true},
		{"overflow", "1e1000", 0.0, true},
		{"underflow", "-1e1000", 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToFloat64(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToFloat64(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToFloat64(%q) unexpected error: %v", tt.input, err)
				}
				// Type assert the result to float64
				floatResult := result.(float64)
				// Use approximate comparison for floating point values
				if math.Abs(floatResult-tt.expected) > 1e-15 {
					t.Errorf("StringToFloat64(%q) = %v, expected %v", tt.input, floatResult, tt.expected)
				}
			}
		})
	}
}
