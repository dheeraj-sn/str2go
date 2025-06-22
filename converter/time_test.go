package converter

import (
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		{
			name:     "RFC3339 format",
			input:    "2023-12-25T15:04:05Z",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC3339Nano format",
			input:    "2023-12-25T15:04:05.123456789Z",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 123456789, time.UTC),
			hasError: false,
		},
		{
			name:     "date time format",
			input:    "2023-12-25 15:04:05",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "date only format",
			input:    "2023-12-25",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC822 format",
			input:    "25 Dec 23 15:04 UTC",
			expected: time.Date(2023, 12, 25, 15, 4, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC850 format",
			input:    "Monday, 25-Dec-23 15:04:05 UTC",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "ANSIC format",
			input:    "Mon Dec 25 15:04:05 2023",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC1123 format",
			input:    "Mon, 25 Dec 2023 15:04:05 UTC",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC1123Z format",
			input:    "Mon, 25 Dec 2023 15:04:05 +0000",
			expected: time.Date(2023, 12, 25, 15, 4, 5, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC822Z format",
			input:    "25 Dec 23 15:04 +0000",
			expected: time.Date(2023, 12, 25, 15, 4, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "invalid format",
			input:    "invalid time format",
			expected: time.Time{},
			hasError: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: time.Time{},
			hasError: true,
		},
		{
			name:     "partial date",
			input:    "2023-12",
			expected: time.Time{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToTime(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("StringToTime(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("StringToTime(%q) unexpected error: %v", tt.input, err)
				}
				timeResult := result.(time.Time)
				if !timeResult.Equal(tt.expected) {
					t.Errorf("StringToTime(%q) = %v, expected %v", tt.input, timeResult.UTC(), tt.expected.UTC())
				}
			}
		})
	}
}
