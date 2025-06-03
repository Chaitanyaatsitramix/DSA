package main

import "fmt"

// ========== VARIABLE SIZE WINDOW METHODS ==========

// Example 1: Longest substring with at most k distinct characters
func LongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}

	left := 0
	maxLen := 0
	charCount := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		// Expand window: add right character
		charCount[s[right]]++

		// Shrink window if more than k distinct characters
		for len(charCount) > k {
			charCount[s[left]]--
			if charCount[s[left]] == 0 {
				delete(charCount, s[left])
			}
			left++
		}

		// Update maximum length
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

// Example 2: Minimum window substring (contains all characters of pattern)
func MinWindowSubstring(s string, pattern string) string {
	if len(s) == 0 || len(pattern) == 0 {
		return ""
	}

	// Count characters in pattern
	patternCount := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		patternCount[pattern[i]]++
	}

	left := 0
	matched := 0
	minLen := len(s) + 1
	minStart := 0

	for right := 0; right < len(s); right++ {
		rightChar := s[right]

		// If right character is part of pattern, decrement count
		if _, exists := patternCount[rightChar]; exists {
			patternCount[rightChar]--
			if patternCount[rightChar] == 0 {
				matched++
			}
		}

		// Shrink window when all characters are matched
		for matched == len(patternCount) {
			// Update minimum window
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			leftChar := s[left]
			if _, exists := patternCount[leftChar]; exists {
				if patternCount[leftChar] == 0 {
					matched--
				}
				patternCount[leftChar]++
			}
			left++
		}
	}

	if minLen > len(s) {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// Example 3: Longest subarray with sum <= target
func LongestSubarrayWithSum(arr []int, target int) int {
	left := 0
	windowSum := 0
	maxLen := 0

	for right := 0; right < len(arr); right++ {
		// Expand window
		windowSum += arr[right]

		// Shrink window while sum exceeds target
		for windowSum > target {
			windowSum -= arr[left]
			left++
		}

		// Update maximum length
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}


// Demo function for Variable Size Window
func DemoVariableWindow() {
	fmt.Println("=== VARIABLE SIZE SLIDING WINDOW EXAMPLES ===")

	// Test Longest Substring K Distinct
	s := "araaci"
	k := 2
	maxLen := LongestSubstringKDistinct(s, k)
	fmt.Printf("Longest substring with at most %d distinct chars in '%s': %d\n", k, s, maxLen)

	// Test Minimum Window Substring
	s2 := "adobecodebanc"
	pattern := "abc"
	minWindow := MinWindowSubstring(s2, pattern)
	fmt.Printf("Minimum window substring of '%s' containing '%s': '%s'\n", s2, pattern, minWindow)

	// Test Longest Subarray With Sum
	arr := []int{2, 1, 2, 3, 4, 1, 2}
	target := 6
	maxLen2 := LongestSubarrayWithSum(arr, target)
	fmt.Printf("Longest subarray with sum <= %d in %v: %d\n", target, arr, maxLen2)
	fmt.Println()
}
