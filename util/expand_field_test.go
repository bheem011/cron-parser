package util

import (
	"reflect"
	"testing"
)

// TestExpandField tests the ExpandField function with multiple test cases
func TestExpandField(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		min, max int
		expected []string
	}{
		{
			name:     "Wildcard",
			field:    "*",
			min:      0,
			max:      59,
			expected: generateRange(0, 59, 1),
		},
		{
			name:     "Step values",
			field:    "*/15",
			min:      0,
			max:      59,
			expected: []string{"0", "15", "30", "45"},
		},
		{
			name:     "Single value",
			field:    "5",
			min:      0,
			max:      59,
			expected: []string{"5"},
		},
		{
			name:     "Range",
			field:    "10-15",
			min:      0,
			max:      59,
			expected: generateRange(10, 15, 1),
		},
		{
			name:     "List of values",
			field:    "5,10,15",
			min:      0,
			max:      59,
			expected: []string{"5", "10", "15"},
		},
		{
			name:     "Range with step",
			field:    "1-12/2",
			min:      0,
			max:      59,
			expected: []string{"1", "3", "5", "7", "9", "11"},
		},
		{
			name:     "Complex list",
			field:    "1-5,10-15/2,20",
			min:      0,
			max:      59,
			expected: []string{"1", "2", "3", "4", "5", "10", "12", "14", "20"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandField(tt.field, tt.min, tt.max)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ExpandField() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestGenerateRange tests the generateRange function
func TestGenerateRange(t *testing.T) {
	tests := []struct {
		name             string
		start, end, step int
		expected         []string
	}{
		{
			name:     "Range 0-5 step 1",
			start:    0,
			end:      5,
			step:     1,
			expected: []string{"0", "1", "2", "3", "4", "5"},
		},
		{
			name:     "Range 1-10 step 2",
			start:    1,
			end:      10,
			step:     2,
			expected: []string{"1", "3", "5", "7", "9"},
		},
		{
			name:     "Range 5-15 step 5",
			start:    5,
			end:      15,
			step:     5,
			expected: []string{"5", "10", "15"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRange(tt.start, tt.end, tt.step)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("generateRange() = %v, want %v", got, tt.expected)
			}
		})
	}
}
