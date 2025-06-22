package converter

import (
	"testing"
)

func TestStringToString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"simple string", "hello", "hello"},
		{"string with spaces", "hello world", "hello world"},
		{"string with special chars", "hello@world#123", "hello@world#123"},
		{"string with unicode", "hello世界", "hello世界"},
		{"string with numbers", "12345", "12345"},
		{"string with mixed content", "Hello123@World!", "Hello123@World!"},
		{"very long string", "this is a very long string with many characters to test the function", "this is a very long string with many characters to test the function"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToString(tt.input)

			if err != nil {
				t.Errorf("StringToString(%q) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("StringToString(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStringToStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{"empty string", "", stringPtr("")},
		{"simple string", "hello", stringPtr("hello")},
		{"string with spaces", "hello world", stringPtr("hello world")},
		{"string with special chars", "hello@world#123", stringPtr("hello@world#123")},
		{"string with unicode", "hello世界", stringPtr("hello世界")},
		{"string with numbers", "12345", stringPtr("12345")},
		{"string with mixed content", "Hello123@World!", stringPtr("Hello123@World!")},
		{"very long string", "this is a very long string with many characters to test the function", stringPtr("this is a very long string with many characters to test the function")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToStringPtr(tt.input)

			if err != nil {
				t.Errorf("StringToStringPtr(%q) unexpected error: %v", tt.input, err)
			}
			if result == nil {
				t.Errorf("StringToStringPtr(%q) expected non-nil result", tt.input)
			} else {
				stringResult := result.(*string)
				if *stringResult != *tt.expected {
					t.Errorf("StringToStringPtr(%q) = %v, expected %v", tt.input, *stringResult, *tt.expected)
				}
			}
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
