package main

import (
	"fmt"
)

// Standard Binary Search - finds exact match
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Prevents overflow

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Target not found
}

// Recursive Binary Search
func BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
}

// Lower Bound - finds first occurrence or insertion point
func LowerBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// Upper Bound - finds position after last occurrence
func UpperBound(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// Find First Occurrence
func FindFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// Find Last Occurrence
func FindLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// Search Range - finds first and last occurrence
func SearchRange(arr []int, target int) [2]int {
	first := FindFirst(arr, target)
	if first == -1 {
		return [2]int{-1, -1}
	}

	last := FindLast(arr, target)
	return [2]int{first, last}
}

// Search in Rotated Sorted Array
func SearchRotated(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		// Check which part is sorted
		if arr[left] <= arr[mid] { // Left part is sorted
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right part is sorted
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// Find Peak Element
func FindPeakElement(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid+1] {
			right = mid // Peak is on the left side or at mid
		} else {
			left = mid + 1 // Peak is on the right side
		}
	}

	return left
}

// Square Root using Binary Search
func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right // Return floor of square root
}

// Search Insert Position
func SearchInsert(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left // Position where target should be inserted
}

// Helper function for demonstration
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test arrays
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	duplicateArray := []int{1, 2, 2, 2, 3, 4, 5}
	rotatedArray := []int{4, 5, 6, 7, 0, 1, 2}
	peakArray := []int{1, 2, 3, 1}

	fmt.Println("=== Binary Search Demonstrations ===")

	// 1. Standard Binary Search
	fmt.Println("1. Standard Binary Search:")
	fmt.Printf("Array: %v\n", sortedArray)
	target := 7
	result := BinarySearch(sortedArray, target)
	fmt.Printf("Searching for %d: Index = %d\n", target, result)

	// Test recursive version
	result2 := BinarySearchRecursive(sortedArray, target, 0, len(sortedArray)-1)
	fmt.Printf("Recursive search for %d: Index = %d\n\n", target, result2)

	// 2. Lower and Upper Bound
	fmt.Println("2. Lower and Upper Bound:")
	fmt.Printf("Array: %v\n", duplicateArray)
	target = 2
	lower := LowerBound(duplicateArray, target)
	upper := UpperBound(duplicateArray, target)
	fmt.Printf("Target %d: Lower Bound = %d, Upper Bound = %d\n\n", target, lower, upper)

	// 3. First and Last Occurrence
	fmt.Println("3. First and Last Occurrence:")
	fmt.Printf("Array: %v\n", duplicateArray)
	searchRange := SearchRange(duplicateArray, target)
	fmt.Printf("Range of %d: [%d, %d]\n\n", target, searchRange[0], searchRange[1])

	// 4. Search in Rotated Array
	fmt.Println("4. Search in Rotated Sorted Array:")
	fmt.Printf("Array: %v\n", rotatedArray)
	target = 0
	result = SearchRotated(rotatedArray, target)
	fmt.Printf("Searching for %d: Index = %d\n\n", target, result)

	// 5. Find Peak Element
	fmt.Println("5. Find Peak Element:")
	fmt.Printf("Array: %v\n", peakArray)
	peak := FindPeakElement(peakArray)
	fmt.Printf("Peak element at index %d: value = %d\n\n", peak, peakArray[peak])

	// 6. Square Root
	fmt.Println("6. Square Root using Binary Search:")
	x := 8
	sqrt := MySqrt(x)
	fmt.Printf("Square root of %d: %d\n\n", x, sqrt)

	// 7. Search Insert Position
	fmt.Println("7. Search Insert Position:")
	fmt.Printf("Array: %v\n", sortedArray)
	target = 6
	insertPos := SearchInsert(sortedArray, target)
	fmt.Printf("Insert position for %d: %d\n", target, insertPos)

	// Demonstration of edge cases
	fmt.Println("\n=== Edge Cases ===")

	// Empty array
	empty := []int{}
	fmt.Printf("Search in empty array: %d\n", BinarySearch(empty, 5))

	// Single element
	single := []int{5}
	fmt.Printf("Search 5 in [5]: %d\n", BinarySearch(single, 5))
	fmt.Printf("Search 3 in [5]: %d\n", BinarySearch(single, 3))

	// Element not found
	fmt.Printf("Search 100 in %v: %d\n", sortedArray, BinarySearch(sortedArray, 100))
}
