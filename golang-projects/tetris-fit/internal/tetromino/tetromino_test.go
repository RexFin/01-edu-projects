package tetromino

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		rawShape []string
		expected []Shape
	}{
		{
			name: "Cube 2x2",
			rawShape: []string{
				"....",
				".##.",
				".##.",
				"....",
			},
			expected: []Shape{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		},
		{
			name: "I shape horizontal",
			rawShape: []string{
				"....",
				"####",
				"....",
				"....",
			},
			expected: []Shape{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
		},
		{
			name: "I shape vertical",
			rawShape: []string{
				"..#.",
				"..#.",
				"..#.",
				"..#.",
			},
			expected: []Shape{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
		},
		{
			name: "T shape",
			rawShape: []string{
				"..#.",
				".##.",
				"..#.",
				"....",
			},
			expected: []Shape{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}},
		},
		{
			name: "L shape",
			rawShape: []string{
				"#...",
				"#...",
				"##..",
				"....",
			},
			expected: []Shape{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}},
		},
		{
			name: "Z shape",
			rawShape: []string{
				"....",
				"##..",
				".##.",
				"....",
			},
			expected: []Shape{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}},
		},
		{
			name: "Empty",
			rawShape: []string{
				"....",
				"....",
				"....",
				"....",
			},
			expected: []Shape{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.rawShape)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("\nwanted: %v\nget:  %v", tt.expected, result)
			}
		})
	}
}
