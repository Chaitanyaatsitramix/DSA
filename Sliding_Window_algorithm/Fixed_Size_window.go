package main

import "fmt"

// ========== FIXED SIZE WINDOW METHODS ==========

// Example 1: Maximum sum of subarray of size k
func MaxSumFixedWindow(arr []int, k int) int {
	if len(arr) < k {
		return -1
	}

	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	maxSum := windowSum

	// Slide the window: remove left, add right
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

// Example 2: Average of all subarrays of size k
func AvgOfSubarrays(arr []int, k int) []float64 {
	if len(arr) < k {
		return []float64{}
	}

	result := []float64{}
	windowSum := 0

	// Calculate sum of first window
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	result = append(result, float64(windowSum)/float64(k))

	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum = windowSum - arr[i-k] + arr[i]
		result = append(result, float64(windowSum)/float64(k))
	}

	return result
}


// Demo function for Fixed Size Window
func DemoFixedWindow() {
	fmt.Println("=== FIXED SIZE SLIDING WINDOW EXAMPLES ===")

	// Test Max Sum Fixed Window
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	maxSum := MaxSumFixedWindow(arr, k)
	fmt.Printf("Max sum of subarray of size %d in %v: %d\n", k, arr, maxSum)

	// Test Average of Subarrays
	arr2 := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k2 := 5
	averages := AvgOfSubarrays(arr2, k2)
	fmt.Printf("Averages of subarrays of size %d: %.2f\n", k2, averages)
	fmt.Println()
}