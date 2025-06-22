package converter

import (
	"testing"
)

func TestStringToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		hasError bool
	}{
		{"true value", "true", true, false},
		{"false value", "false", false, false},
		{"TRUE value", "TRUE", true, false},
		{"FALSE value", "FALSE", false, false},
		{"True value", "True", true, false},
		{"False value", "False", false, false},
		{"1 value", "1", true, false},
		{"0 value", "0", false, false},
		{"invalid value", "invalid", false, true},
		{"empty string", "", false, true},
		{"yes value", "yes", false, true},
		{"no value", "no", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToBool(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToBool(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToBool(%q) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("StringToBool(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestStringToBoolPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *bool
		hasError bool
	}{
		{"true value", "true", boolPtr(true), false},
		{"false value", "false", boolPtr(false), false},
		{"TRUE value", "TRUE", boolPtr(true), false},
		{"FALSE value", "FALSE", boolPtr(false), false},
		{"True value", "True", boolPtr(true), false},
		{"False value", "False", boolPtr(false), false},
		{"1 value", "1", boolPtr(true), false},
		{"0 value", "0", boolPtr(false), false},
		{"invalid value", "invalid", nil, true},
		{"empty string", "", nil, true},
		{"yes value", "yes", nil, true},
		{"no value", "no", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToBoolPtr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToBoolPtr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToBoolPtr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToBoolPtr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToBoolPtr(%q) expected non-nil result", tt.input)
				} else {
					boolResult := result.(*bool)
					if *boolResult != *tt.expected {
						t.Errorf("StringToBoolPtr(%q) = %v, expected %v", tt.input, *boolResult, *tt.expected)
					}
				}
			}
		})
	}
}

// Helper function to create bool pointers
func boolPtr(b bool) *bool {
	return &b
}
