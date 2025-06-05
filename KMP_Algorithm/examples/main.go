package main

import (
	"fmt"

	"KMP_Algorithm"
)

func main() {
	fmt.Println("=== KMP (Knuth-Morris-Pratt) Algorithm Demonstrations ===\n")

	// Example 1: Basic Pattern Matching
	text := "AABAACAADAABAAABAA"
	pattern := "AABA"
	fmt.Printf("Text: %s\nPattern: %s\n", text, pattern)

	// Show the prefix table (LPS array)
	lps := kmp.ComputeLPSArray(pattern)
	fmt.Printf("Prefix Table (LPS): %v\n", lps)

	// Find all occurrences
	matches := kmp.KMPSearch(text, pattern)
	fmt.Printf("All matches found at indices: %v\n\n", matches)

	// Example 2: First Occurrence Only
	text2 := "ABABDABACDABABCABAB"
	pattern2 := "ABABCABAB"
	fmt.Printf("Text: %s\nPattern: %s\n", text2, pattern2)
	firstMatch := kmp.KMPSearchFirst(text2, pattern2)
	fmt.Printf("First match found at index: %d\n\n", firstMatch)

	// Example 3: Overlapping Patterns
	text3 := "AAAAA"
	pattern3 := "AA"
	fmt.Printf("Text: %s\nPattern: %s\n", text3, pattern3)

	// Normal search (non-overlapping)
	normalMatches := kmp.KMPSearch(text3, pattern3)
	fmt.Printf("Non-overlapping matches at indices: %v\n", normalMatches)

	// Overlapping search
	overlapMatches := kmp.KMPSearchWithOverlap(text3, pattern3)
	fmt.Printf("Overlapping matches at indices: %v\n\n", overlapMatches)

	// Example 4: Pattern Count
	text4 := "AAAAABAAABA"
	pattern4 := "AA"
	fmt.Printf("Text: %s\nPattern: %s\n", text4, pattern4)
	count := kmp.KMPCount(text4, pattern4)
	fmt.Printf("Pattern occurs %d times\n\n", count)

	// Example 5: Edge Cases
	fmt.Println("Edge Cases:")
	fmt.Printf("Empty text search: %v\n", kmp.KMPSearch("", "ABC"))
	fmt.Printf("Empty pattern search: %v\n", kmp.KMPSearch("ABC", ""))
	fmt.Printf("Pattern longer than text: %d\n", kmp.KMPSearchFirst("AB", "ABC"))
}
