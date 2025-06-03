package main

import "fmt"

// ========== CLASSIC KADANE'S ALGORITHM ==========

// Example 1: Basic Kadane's Algorithm - Maximum Sum Subarray
func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSoFar := arr[0]      // Maximum sum found so far
	maxEndingHere := arr[0] // Maximum sum ending at current position

	for i := 1; i < len(arr); i++ {
		// Either start new subarray from current element or extend existing one
		maxEndingHere = max(arr[i], maxEndingHere+arr[i])

		// Update global maximum
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ========== KADANE'S WITH INDICES ==========

// Example 2: Kadane's Algorithm with start and end indices
func MaxSubarrayWithIndices(arr []int) (int, int, int) {
	if len(arr) == 0 {
		return 0, -1, -1
	}

	maxSoFar := arr[0]
	maxEndingHere := arr[0]
	start := 0
	end := 0
	tempStart := 0

	for i := 1; i < len(arr); i++ {
		// If starting new subarray is better
		if arr[i] > maxEndingHere+arr[i] {
			maxEndingHere = arr[i]
			tempStart = i
		} else {
			maxEndingHere = maxEndingHere + arr[i]
		}

		// Update global maximum and indices
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
			start = tempStart
			end = i
		}
	}

	return maxSoFar, start, end
}

// ========== KADANE'S FOR CIRCULAR ARRAY ==========

// Example 3: Maximum Sum Circular Subarray
func MaxCircularSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// Case 1: Maximum subarray is non-circular (normal Kadane's)
	maxKadane := MaxSubarraySum(arr)

	// Case 2: Maximum subarray is circular
	// This means we need to find minimum subarray and subtract from total
	arraySum := 0
	for i := 0; i < len(arr); i++ {
		arraySum += arr[i]
		arr[i] = -arr[i] // Negate for finding minimum
	}

	// Find minimum subarray sum (which is max of negated array)
	minSubarraySum := -MaxSubarraySum(arr)

	// Restore original array
	for i := 0; i < len(arr); i++ {
		arr[i] = -arr[i]
	}

	// Maximum circular sum = Total sum - Minimum subarray sum
	maxCircular := arraySum - minSubarraySum

	// Handle edge case: if all elements are negative
	if maxCircular == 0 {
		return maxKadane
	}

	return max(maxKadane, maxCircular)
}

// ========== VARIATIONS OF KADANE'S ==========

// Example 4: Maximum Product Subarray
func MaxProductSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSoFar := arr[0]
	maxEndingHere := arr[0]
	minEndingHere := arr[0]

	for i := 1; i < len(arr); i++ {
		// Store the previous max as we'll update it
		temp := maxEndingHere

		// Maximum product ending here is max of:
		// 1. Current element
		// 2. Current element * previous max
		// 3. Current element * previous min (in case current is negative)
		maxEndingHere = max(arr[i], max(arr[i]*maxEndingHere, arr[i]*minEndingHere))

		// Minimum product ending here
		minEndingHere = min(arr[i], min(arr[i]*temp, arr[i]*minEndingHere))

		// Update global maximum
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// Example 5: Maximum Sum with No Three Consecutive Elements
func MaxSumNoThreeConsecutive(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	if n == 2 {
		return arr[0] + arr[1]
	}

	// dp[i] represents maximum sum up to index i
	dp := make([]int, n)
	dp[0] = arr[0]
	dp[1] = arr[0] + arr[1]
	dp[2] = max(dp[1], max(arr[1]+arr[2], arr[0]+arr[2]))

	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], max(dp[i-2]+arr[i], dp[i-3]+arr[i-1]+arr[i]))
	}

	return dp[n-1]
}

// ========== DEMO FUNCTIONS ==========

// Demo function for Classic Kadane's
func DemoClassicKadane() {
	fmt.Println("=== CLASSIC KADANE'S ALGORITHM ===")

	// Test cases
	testArrays := [][]int{
		{-2, -3, 4, -1, -2, 1, 5, -3},
		{1, 2, 3, 4, 5},
		{-1, -2, -3, -4},
		{5},
		{-5, -2, -8, -1},
	}

	for _, arr := range testArrays {
		maxSum := MaxSubarraySum(arr)
		fmt.Printf("Array: %v\n", arr)
		fmt.Printf("Maximum subarray sum: %d\n\n", maxSum)
	}
}

// Demo function for Kadane's with indices
func DemoKadaneWithIndices() {
	fmt.Println("=== KADANE'S WITH INDICES ===")

	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	maxSum, start, end := MaxSubarrayWithIndices(arr)

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Maximum subarray sum: %d\n", maxSum)
	fmt.Printf("Subarray indices: [%d, %d]\n", start, end)
	fmt.Printf("Subarray elements: %v\n\n", arr[start:end+1])
}

// Demo function for Circular Kadane's
func DemoCircularKadane() {
	fmt.Println("=== CIRCULAR KADANE'S ALGORITHM ===")

	testArrays := [][]int{
		{10, -3, -4, 7, 6, 5, -4, -1},
		{-3, -2, -3},
		{3, -2, 2, -3},
		{5, -3, 5},
	}

	for _, arr := range testArrays {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)

		maxSum := MaxCircularSubarray(arrCopy)
		fmt.Printf("Array: %v\n", arr)
		fmt.Printf("Maximum circular subarray sum: %d\n\n", maxSum)
	}
}

// Demo function for variations
func DemoKadaneVariations() {
	fmt.Println("=== KADANE'S ALGORITHM VARIATIONS ===")

	// Maximum Product Subarray
	arr1 := []int{2, 3, -2, 4}
	maxProduct := MaxProductSubarray(arr1)
	fmt.Printf("Maximum Product Subarray in %v: %d\n", arr1, maxProduct)

	// Maximum Sum with No Three Consecutive
	arr2 := []int{1, 2, 3}
	maxSumNoThree := MaxSumNoThreeConsecutive(arr2)
	fmt.Printf("Max sum with no three consecutive in %v: %d\n", arr2, maxSumNoThree)
	fmt.Println()
}
