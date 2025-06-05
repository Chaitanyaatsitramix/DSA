package main

import (
	"fmt"
)

// QuickSelect finds the kth smallest element in an unordered array
func QuickSelect(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}
	return quickSelectHelper(arr, 0, len(arr)-1, k)
}

// quickSelectHelper is the recursive helper function for QuickSelect
func quickSelectHelper(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}
	pivotIdx := partition(arr, left, right)
	if k == pivotIdx {
		return arr[k]
	}
	if k < pivotIdx {
		return quickSelectHelper(arr, left, pivotIdx-1, k)
	}
	return quickSelectHelper(arr, pivotIdx+1, right, k)
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func main() {
	fmt.Println("=== Quickselect Algorithm Demonstrations ===\n")

	// Example 1: Finding kth smallest element
	arr1 := []int{7, 10, 4, 3, 20, 15}
	k := 3 // Find the 4th smallest element (0-based index)
	fmt.Printf("Array: %v\n", arr1)
	fmt.Printf("Finding %dth smallest element (k=%d)\n", k+1, k)
	result := QuickSelect(arr1, k)
	fmt.Printf("Result: %d\n\n", result)

	// Example 2: Finding median
	arr2 := []int{12, 3, 5, 7, 4, 19, 26}
	medianIndex := len(arr2) / 2
	fmt.Printf("Array: %v\n", arr2)
	fmt.Printf("Finding median (k=%d)\n", medianIndex)
	median := QuickSelect(arr2, medianIndex)
	fmt.Printf("Median: %d\n\n", median)

	// Example 3: Edge cases
	arr4 := []int{1} // Single element
	fmt.Printf("Single element array: %v\n", arr4)
	fmt.Printf("Finding smallest element (k=0)\n")
	result = QuickSelect(arr4, 0)
	fmt.Printf("Result: %d\n\n", result)

	// Example 4: Duplicate elements
	arr5 := []int{3, 3, 3, 3, 3, 3, 3}
	k = 3
	fmt.Printf("Array with duplicates: %v\n", arr5)
	fmt.Printf("Finding %dth element (k=%d)\n", k+1, k)
	result = QuickSelect(arr5, k)
	fmt.Printf("Result: %d\n", result)
}
