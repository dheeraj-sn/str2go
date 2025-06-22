package converter

import (
	"testing"
)

func TestStringToUint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
		hasError bool
	}{
		{"zero", "0", 0, false},
		{"positive integer", "123", 123, false},
		{"max uint64", "18446744073709551615", 18446744073709551615, false},
		{"negative value", "-1", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
		{"overflow", "18446744073709551616", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint(%q) unexpected error: %v", tt.input, err)
				}
				if result.(uint64) != tt.expected {
					t.Errorf("StringToUint(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestStringToUint8(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint8
		hasError bool
	}{
		{"zero", "0", 0, false},
		{"positive integer", "123", 123, false},
		{"max uint8", "255", 255, false},
		{"overflow", "256", 0, true},
		{"negative value", "-1", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint8(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint8(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint8(%q) unexpected error: %v", tt.input, err)
				}
				if result.(uint64) != uint64(tt.expected) {
					t.Errorf("StringToUint8(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestStringToUint16(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint16
		hasError bool
	}{
		{"zero", "0", 0, false},
		{"positive integer", "12345", 12345, false},
		{"max uint16", "65535", 65535, false},
		{"overflow", "65536", 0, true},
		{"negative value", "-1", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint16(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint16(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint16(%q) unexpected error: %v", tt.input, err)
				}
				if result.(uint64) != uint64(tt.expected) {
					t.Errorf("StringToUint16(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestStringToUint32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint32
		hasError bool
	}{
		{"zero", "0", 0, false},
		{"positive integer", "123456", 123456, false},
		{"max uint32", "4294967295", 4294967295, false},
		{"overflow", "4294967296", 0, true},
		{"negative value", "-1", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint32(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint32(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint32(%q) unexpected error: %v", tt.input, err)
				}
				if result.(uint64) != uint64(tt.expected) {
					t.Errorf("StringToUint32(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestStringToUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
		hasError bool
	}{
		{"zero", "0", 0, false},
		{"positive integer", "123456789012345", 123456789012345, false},
		{"max uint64", "18446744073709551615", 18446744073709551615, false},
		{"overflow", "18446744073709551616", 0, true},
		{"negative value", "-1", 0, true},
		{"decimal", "123.456", 0, true},
		{"invalid string", "invalid", 0, true},
		{"empty string", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint64(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint64(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint64(%q) unexpected error: %v", tt.input, err)
				}
				if result.(uint64) != tt.expected {
					t.Errorf("StringToUint64(%q) = %v, expected %v", tt.input, result, tt.expected)
				}
			}
		})
	}
}
