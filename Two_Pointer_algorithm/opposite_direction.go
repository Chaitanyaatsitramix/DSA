package main

import "fmt"

// Example 1: Two Sum - Find pair that adds up to target
func TwoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++ // Need bigger sum
		} else {
			right-- // Need smaller sum
		}
	}
	return []int{-1, -1} // Not found
}

// Example 2: Check if string is palindrome
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Example 3: Reverse array in-place
func ReverseArray(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// Demo function to show examples
func DemoOppositeDirection() {
	fmt.Println("=== OPPOSITE DIRECTION (Convergent) EXAMPLES ===")

	// Test Two Sum (assuming sorted array)
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Printf("Two Sum: indices %v sum to %d\n", result, target)

	// Test Palindrome
	word := "racecar"
	fmt.Printf("'%s' is palindrome: %v\n", word, IsPalindrome(word))

	// Test Reverse Array
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", arr)
	ReverseArray(arr)
	fmt.Printf("Reversed: %v\n", arr)
	fmt.Println()
}
