package converter

import (
	"testing"
	"time"
)

func TestStringToTimePtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *time.Time
		hasError bool
	}{
		{
			name:     "RFC3339 format",
			input:    "2023-12-25T15:04:05Z",
			expected: timePtr(time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC)),
			hasError: false,
		},
		{
			name:     "RFC3339 with timezone",
			input:    "2023-12-25T15:04:05-07:00",
			expected: timePtr(time.Date(2023, 12, 25, 15, 4, 5, 0, time.FixedZone("", -7*3600))),
			hasError: false,
		},
		{
			name:     "RFC3339 with nanoseconds",
			input:    "2023-12-25T15:04:05.123456789Z",
			expected: timePtr(time.Date(2023, 12, 25, 15, 4, 5, 123456789, time.UTC)),
			hasError: false,
		},
		{
			name:     "invalid format",
			input:    "invalid time format",
			expected: nil,
			hasError: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
			hasError: true,
		},
		{
			name:     "partial date",
			input:    "2023-12",
			expected: nil,
			hasError: true,
		},
		{
			name:     "wrong format",
			input:    "2023-12-25 15:04:05",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToTimePtr(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToTimePtr(%q) expected error, got nil", tt.input)
				}
				if result != nil {
					t.Errorf("StringToTimePtr(%q) expected nil result, got %v", tt.input, result)
				}
			} else {
				if err != nil {
					t.Errorf("StringToTimePtr(%q) unexpected error: %v", tt.input, err)
				}
				if result == nil {
					t.Errorf("StringToTimePtr(%q) expected non-nil result", tt.input)
				} else {
					timeResult := result.(*time.Time)
					if timeResult == nil {
						t.Errorf("StringToTimePtr(%q) expected non-nil time pointer", tt.input)
					} else {
						// Compare only the date and time components, not the location
						if !timeResult.Equal(*tt.expected) {
							t.Errorf("StringToTimePtr(%q) = %v, expected %v", tt.input, *timeResult, *tt.expected)
						}
					}
				}
			}
		})
	}
}

// Helper function to create time pointers
func timePtr(t time.Time) *time.Time {
	return &t
}
