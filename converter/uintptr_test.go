package converter

import (
	"testing"
)

func TestStringToUintPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *uint64
		hasError bool
	}{
		{"zero", "0", uint64Ptr(0), false},
		{"positive integer", "123", uint64Ptr(123), false},
		{"max uint64", "18446744073709551615", uint64Ptr(18446744073709551615), false},
		{"negative value", "-1", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
		{"overflow", "18446744073709551616", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUintPtr(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUintPtr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToUintPtr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUintPtr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToUintPtr(%q) expected non-nil result", tt.input)
				} else {
					uintResult := result.(*uint64)
					if *uintResult != *tt.expected {
						t.Errorf("StringToUintPtr(%q) = %v, expected %v", tt.input, *uintResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToUint8Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *uint8
		hasError bool
	}{
		{"zero", "0", uint8Ptr(0), false},
		{"positive integer", "123", uint8Ptr(123), false},
		{"max uint8", "255", uint8Ptr(255), false},
		{"overflow", "256", nil, true},
		{"negative value", "-1", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint8Ptr(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint8Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToUint8Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint8Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToUint8Ptr(%q) expected non-nil result", tt.input)
				} else {
					uintResult := result.(*uint64)
					if *uintResult != uint64(*tt.expected) {
						t.Errorf("StringToUint8Ptr(%q) = %v, expected %v", tt.input, *uintResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToUint16Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *uint16
		hasError bool
	}{
		{"zero", "0", uint16Ptr(0), false},
		{"positive integer", "12345", uint16Ptr(12345), false},
		{"max uint16", "65535", uint16Ptr(65535), false},
		{"overflow", "65536", nil, true},
		{"negative value", "-1", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint16Ptr(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint16Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToUint16Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint16Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToUint16Ptr(%q) expected non-nil result", tt.input)
				} else {
					uintResult := result.(*uint64)
					if *uintResult != uint64(*tt.expected) {
						t.Errorf("StringToUint16Ptr(%q) = %v, expected %v", tt.input, *uintResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToUint32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *uint32
		hasError bool
	}{
		{"zero", "0", uint32Ptr(0), false},
		{"positive integer", "123456", uint32Ptr(123456), false},
		{"max uint32", "4294967295", uint32Ptr(4294967295), false},
		{"overflow", "4294967296", nil, true},
		{"negative value", "-1", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint32Ptr(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint32Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToUint32Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint32Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToUint32Ptr(%q) expected non-nil result", tt.input)
				} else {
					uintResult := result.(*uint64)
					if *uintResult != uint64(*tt.expected) {
						t.Errorf("StringToUint32Ptr(%q) = %v, expected %v", tt.input, *uintResult, *tt.expected)
					}
				}
			}
		})
	}
}

func TestStringToUint64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *uint64
		hasError bool
	}{
		{"zero", "0", uint64Ptr(0), false},
		{"positive integer", "123456789012345", uint64Ptr(123456789012345), false},
		{"max uint64", "18446744073709551615", uint64Ptr(18446744073709551615), false},
		{"overflow", "18446744073709551616", nil, true},
		{"negative value", "-1", nil, true},
		{"decimal", "123.456", nil, true},
		{"invalid string", "invalid", nil, true},
		{"empty string", "", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToUint64Ptr(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("StringToUint64Ptr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToUint64Ptr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToUint64Ptr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToUint64Ptr(%q) expected non-nil result", tt.input)
				} else {
					uintResult := result.(*uint64)
					if *uintResult != *tt.expected {
						t.Errorf("StringToUint64Ptr(%q) = %v, expected %v", tt.input, *uintResult, *tt.expected)
					}
				}
			}
		})
	}
}

// Helper functions to create uint pointers
func uint8Ptr(u uint8) *uint8 {
	return &u
}
func uint16Ptr(u uint16) *uint16 {
	return &u
}
func uint32Ptr(u uint32) *uint32 {
	return &u
}
func uint64Ptr(u uint64) *uint64 {
	return &u
}
