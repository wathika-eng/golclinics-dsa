package main

import (
	"reflect"
	"testing"
)

// TestSortedArr tests the sortedArr function
func TestSortedArr(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: []int{10, -3, 2, 1, 4}, output: []int{-3, 1, 2, 4, 10}},
		{input: []int{5, 1, 3, 2, 4}, output: []int{1, 2, 3, 4, 5}},
		{input: []int{}, output: []int{}},
		{input: []int{1}, output: []int{1}},
		{input: []int{-1, -5, 0}, output: []int{-5, -1, 0}},
	}

	for _, test := range tests {
		result := sortedArr(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("sortedArr(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}

// TestReversedArr tests the reversedArr function
func TestReversedArr(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{input: []int{10, -3, 2, 1, 4}, output: []int{4, 1, 2, -3, 10}},
		{input: []int{1, 2, 3, 4, 5}, output: []int{5, 4, 3, 2, 1}},
		{input: []int{}, output: []int{}},
		{input: []int{1}, output: []int{1}},
	}

	for _, test := range tests {
		result := reversedArr(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("reversedArr(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}
