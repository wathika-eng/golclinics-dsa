package main

import (
	"fmt"
	"slices"
)

func sortedArr(A []int) []int {
	// get the length of the array
	length := len(A)
	// transverse the whole array
	for i := 0; i < length; i++ {
		// smallest number is the first index or 0
		minIndex := i
		// inner loop starts from i + 1
		for j := i + 1; j < length; j++ {
			// update minvalue to smaller value if found
			if A[j] < A[minIndex] {
				// min value
				minIndex = j
			}
		}
		// if minimum value is not i, swap it
		if minIndex != i {
			temp := A[minIndex]
			A[minIndex] = A[i]
			A[i] = temp
		}
	}
	return A
}
func reversedArr(A []int) []int {
	slices.Reverse(A)
  return A
	// length := len(A)
	// for i := length - 1; i >= 0; i--{
	//   slices.Reverse(A)
	// }
}
func main() {
	array := []int{10, -3, 2, 1, 4}

	fmt.Println("Sort the Array: ", sortedArr(array))
	fmt.Println("Reverse the Array: ", reversedArr(array))
}
