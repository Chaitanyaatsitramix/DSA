package kmp

// ComputeLPSArray builds the Longest Proper Prefix which is also Suffix array
// This is the core of KMP Algorithm that enables O(n) string matching
func ComputeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)

	// Length of the previous longest prefix suffix
	length := 0
	i := 1

	// First value is always 0
	lps[0] = 0

	// Calculate lps[i] for i = 1 to m-1
	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				// Try to find a shorter prefix which is also a suffix
				length = lps[length-1]
			} else {
				// No proper prefix found
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// KMPSearch finds all occurrences of pattern in text using KMP algorithm
// Returns an array of starting indices where pattern is found
func KMPSearch(text, pattern string) []int {
	matches := []int{}
	if len(pattern) == 0 || len(text) == 0 {
		return matches
	}

	// Get the prefix table
	lps := ComputeLPSArray(pattern)

	i := 0 // Index for text
	j := 0 // Index for pattern

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			// Pattern found at index i-j
			matches = append(matches, i-j)
			// Look for the next match
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches
}

// KMPSearchFirst finds the first occurrence of pattern in text
// Returns the index of first match, or -1 if not found
func KMPSearchFirst(text, pattern string) int {
	if len(pattern) == 0 || len(text) == 0 {
		return -1
	}

	lps := ComputeLPSArray(pattern)
	i, j := 0, 0

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			return i - j
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

// KMPCount counts the number of occurrences of pattern in text
func KMPCount(text, pattern string) int {
	return len(KMPSearch(text, pattern))
}

// KMPSearchWithOverlap finds all occurrences including overlapping matches
func KMPSearchWithOverlap(text, pattern string) []int {
	matches := []int{}
	if len(pattern) == 0 || len(text) == 0 {
		return matches
	}

	lps := ComputeLPSArray(pattern)
	i, j := 0, 0

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			// For overlapping matches, only move j back one position
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches
}
