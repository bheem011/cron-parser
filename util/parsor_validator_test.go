package util

import (
	"testing"
)

// TestValidCornParser tests the ValidCornParser function with both valid and invalid cron strings
func TestValidCornParser(t *testing.T) {
	tests := []struct {
		name     string
		cron     string
		expected bool
	}{
		// Valid cron strings
		{
			name:     "Valid wildcard",
			cron:     "* * * * * /usr/bin/find",
			expected: true,
		},
		{
			name:     "Valid step values",
			cron:     "*/15 0 1,15 * 1-5 /usr/bin/find",
			expected: true,
		},
		{
			name:     "Valid single value",
			cron:     "0 12 * * * /usr/bin/backup",
			expected: true,
		},
		{
			name:     "Valid range",
			cron:     "0 0-23/2 * * * /usr/bin/find",
			expected: true,
		},
		{
			name:     "Valid complex",
			cron:     "0 0 1-15,20-25/2 * 1-5 /usr/bin/find",
			expected: true,
		},
		{
			name:     "Valid month and day",
			cron:     "0 12 1-5 1-12 1-5 /usr/bin/find",
			expected: true,
		},
		{
			name:     "Valid range with step",
			cron:     "0 0 1-12/2 * * /usr/bin/find",
			expected: true,
		},

		// Invalid cron strings
		{
			name:     "Invalid missing command",
			cron:     "*/5 0 1,15 * 1-5",
			expected: false,
		},
		{
			name:     "Invalid extra space",
			cron:     "*/5 0 1,15 * 1-5  /usr/bin/find",
			expected: false,
		},
		{
			name:     "Invalid invalid step",
			cron:     "*/70 * * * * /usr/bin/find",
			expected: false,
		},
		{
			name:     "Invalid out of range hour",
			cron:     "0 25 * * * /usr/bin/find",
			expected: false,
		},
		{
			name:     "Invalid non-numeric",
			cron:     "a b c d e /usr/bin/find",
			expected: false,
		},
		{
			name:     "Invalid step without range",
			cron:     "*/15 */25 1,15 * 1-5 /usr/bin/find",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidCornParser(tt.cron)
			if got != tt.expected {
				t.Errorf("ValidCornParser() = %v, want %v for cron string %q", got, tt.expected, tt.cron)
			}
		})
	}
}
