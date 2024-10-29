package main

import (
	"fmt"
	"testing"
)

func TestBasicLinearSearch(t *testing.T) {
	tests := []struct {
		haystack []int
		needle   int
		expected []int
		found    bool
	}{
		{haystack: []int{3, 2, 4}, needle: 6, expected: []int{1, 2}, found: true},
		{haystack: []int{1, 2, 3, 4}, needle: 5, expected: []int{0, 3}, found: true},
		{haystack: []int{5, 1, 2, 3}, needle: 8, expected: []int{0, 3}, found: true},
		{haystack: []int{1, 2, 3}, needle: 7, expected: []int{}, found: false},
		{haystack: []int{0, 0, 0}, needle: 0, expected: []int{0, 1}, found: true}, // Multiple zeros
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%d", tt.haystack, tt.needle), func(t *testing.T) {
			gotFound, gotIndices := basicLinearSearch(tt.haystack, tt.needle)
			if gotFound != tt.found {
				t.Errorf("Expected found: %v, got: %v", tt.found, gotFound)
			}
			if (gotFound && len(gotIndices) != 2) || (!gotFound && len(gotIndices) != 0) {
				t.Errorf("Expected indices length: %v, got: %v", len(tt.expected), len(gotIndices))
			}
			if gotFound && len(gotIndices) == 2 {
				if gotIndices[0] != tt.expected[0] || gotIndices[1] != tt.expected[1] {
					t.Errorf("Expected indices: %v, got: %v", tt.expected, gotIndices)
				}
			}
		})
	}
}
