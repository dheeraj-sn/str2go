package converter

import (
	"math"
	"testing"
)

func TestStringToFloat32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *float32
		hasError bool
	}{
		{"positive integer", "123", float32Ptr(123.0), false},
		{"negative integer", "-123", float32Ptr(-123.0), false},
		{"positive decimal", "123.456", float32Ptr(123.456), false},
		{"negative decimal", "-123.456", float32Ptr(-123.456), false},
		{"zero", "0", float32Ptr(0.0), false},
		{"zero decimal", "0.0", float32Ptr(0.0), false},
		{"scientific notation", "1.23e2", float32Ptr(123.0), false},
		{"negative scientific notation", "-1.23e2", float32Ptr(-123.0), false},
		{"large number", "999999.999", float32Ptr(999999.999), false},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
		{"overflow", "1e100", nil, true},
		{"underflow", "-1e100", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToFloat32Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToFloat32Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToFloat32Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToFloat32Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToFloat32Ptr(%q) expected non-nil result", tt.input)
				} else {
					// Type assert the result to *float64 (since ParseFloat returns float64)
					floatResult := result.(*float64)
					if floatResult == nil {
						t.Errorf("StringToFloat32Ptr(%q) expected non-nil float pointer", tt.input)
					} else {
						// Use approximate comparison for floating point values
						if math.Abs(*floatResult-float64(*tt.expected)) > 1e-6 {
							t.Errorf("StringToFloat32Ptr(%q) = %v, expected %v", tt.input, *floatResult, *tt.expected)
						}
					}
				}
			}
		})
	}
}

func TestStringToFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *float64
		hasError bool
	}{
		{"positive integer", "123", float64Ptr(123.0), false},
		{"negative integer", "-123", float64Ptr(-123.0), false},
		{"positive decimal", "123.456", float64Ptr(123.456), false},
		{"negative decimal", "-123.456", float64Ptr(-123.456), false},
		{"zero", "0", float64Ptr(0.0), false},
		{"zero decimal", "0.0", float64Ptr(0.0), false},
		{"scientific notation", "1.23e2", float64Ptr(123.0), false},
		{"negative scientific notation", "-1.23e2", float64Ptr(-123.0), false},
		{"large number", "999999.999", float64Ptr(999999.999), false},
		{"very large number", "1e308", float64Ptr(1e308), false},
		{"very small number", "1e-308", float64Ptr(1e-308), false},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
		{"overflow", "1e1000", nil, true},
		{"underflow", "-1e1000", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToFloat64Ptr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToFloat64Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToFloat64Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToFloat64Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToFloat64Ptr(%q) expected non-nil result", tt.input)
				} else {
					// Type assert the result to *float64
					floatResult := result.(*float64)
					if floatResult == nil {
						t.Errorf("StringToFloat64Ptr(%q) expected non-nil float pointer", tt.input)
					} else {
						// Use approximate comparison for floating point values
						if math.Abs(*floatResult-*tt.expected) > 1e-15 {
							t.Errorf("StringToFloat64Ptr(%q) = %v, expected %v", tt.input, *floatResult, *tt.expected)
						}
					}
				}
			}
		})
	}
}

// Helper functions to create float pointers
func float32Ptr(f float32) *float32 {
	return &f
}

func float64Ptr(f float64) *float64 {
	return &f
}
