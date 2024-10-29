package main

import "fmt"

func basicLinearSearch(haystack []int, needle int) (bool, []int) {
	length := len(haystack)
	for i := range haystack {
		for j := i + 1; j < length; j++ {
			if haystack[i]+haystack[j] == needle {
				return true, []int{i, j}
			}
		}

	}
	return false, []int{}
}

func main() {
	haystack := []int{3, 2, 4}
	needle := 6
	fmt.Println(basicLinearSearch(haystack, needle))
}

// arr := make([]int, 0, 10)
// // arr[0] = 1
// fmt.Println("Size in bytes:", unsafe.Sizeof(arr)) // Size of the entire array in bytes
// fmt.Println("Length of array:", len(arr))         // Number of elements in the array
// fmt.Println(arr)
