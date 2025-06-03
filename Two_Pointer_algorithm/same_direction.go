package main

import "fmt"

// Example 1: Remove duplicates from sorted array
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0 // Points to position of unique elements
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1 // Length of unique elements
}

// Example 2: Find maximum sum of subarray of size k
func MaxSumSubarray(arr []int, k int) int {
	if len(arr) < k {
		return -1
	}

	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i] // Remove left, add right
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// Demo function to show examples
func DemoSameDirection() {
	fmt.Println("=== SAME DIRECTION (Sliding Window) EXAMPLES ===")

	// Test Remove Duplicates
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Printf("Original: %v\n", nums)
	length := RemoveDuplicates(nums)
	fmt.Printf("After removing duplicates: %v (length: %d)\n", nums[:length], length)

	// Test Max Sum Subarray
	arr := []int{1,2,3,4,5}
	k := 3
	maxSum := MaxSumSubarray(arr, k)
	fmt.Printf("Max sum of subarray of size %d: %d\n", k, maxSum)

}
