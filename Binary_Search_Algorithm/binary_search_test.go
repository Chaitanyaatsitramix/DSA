package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}

	tests := []struct {
		target   int
		expected int
	}{
		{7, 3},
		{1, 0},
		{13, 6},
		{4, -1},  // Not found
		{15, -1}, // Greater than all
		{0, -1},  // Less than all
	}

	for _, test := range tests {
		result := BinarySearch(arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10}

	tests := []struct {
		target   int
		expected int
	}{
		{6, 2},
		{2, 0},
		{10, 4},
		{5, -1},
		{12, -1},
	}

	for _, test := range tests {
		result := BinarySearchRecursive(arr, test.target, 0, len(arr)-1)
		if result != test.expected {
			t.Errorf("BinarySearchRecursive(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestLowerBound(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 5}

	tests := []struct {
		target   int
		expected int
	}{
		{2, 1}, // First occurrence of 2
		{3, 4}, // First occurrence of 3
		{0, 0}, // Before all elements
		{6, 7}, // After all elements
		{2, 1}, // Multiple occurrences
	}

	for _, test := range tests {
		result := LowerBound(arr, test.target)
		if result != test.expected {
			t.Errorf("LowerBound(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestUpperBound(t *testing.T) {
	arr := []int{1, 2, 2, 2, 3, 4, 5}

	tests := []struct {
		target   int
		expected int
	}{
		{2, 4}, // After last occurrence of 2
		{3, 5}, // After last occurrence of 3
		{0, 0}, // Before all elements
		{6, 7}, // After all elements
		{1, 1}, // Single occurrence
	}

	for _, test := range tests {
		result := UpperBound(arr, test.target)
		if result != test.expected {
			t.Errorf("UpperBound(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestSearchRange(t *testing.T) {
	arr := []int{5, 7, 7, 8, 8, 10}

	tests := []struct {
		target   int
		expected [2]int
	}{
		{8, [2]int{3, 4}},   // Range of 8
		{7, [2]int{1, 2}},   // Range of 7
		{5, [2]int{0, 0}},   // Single occurrence
		{6, [2]int{-1, -1}}, // Not found
		{10, [2]int{5, 5}},  // Single occurrence at end
	}

	for _, test := range tests {
		result := SearchRange(arr, test.target)
		if result != test.expected {
			t.Errorf("SearchRange(%v, %d) = %v; expected %v", arr, test.target, result, test.expected)
		}
	}
}

func TestSearchRotated(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 4},
		{3, -1}, // Not found
		{4, 0},  // First element
		{2, 6},  // Last element
		{7, 3},  // In first part
		{1, 5},  // In second part
	}

	for _, test := range tests {
		result := SearchRotated(arr, test.target)
		if result != test.expected {
			t.Errorf("SearchRotated(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool // Just check if result is valid peak
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 1, 3, 5, 6, 4}, true},
		{[]int{1}, true},    // Single element
		{[]int{1, 2}, true}, // Two elements
		{[]int{2, 1}, true}, // Two elements descending
	}

	for _, test := range tests {
		result := FindPeakElement(test.arr)
		isPeak := false

		// Check if result is a valid peak
		if len(test.arr) == 1 {
			isPeak = (result == 0)
		} else if result == 0 {
			isPeak = test.arr[0] > test.arr[1]
		} else if result == len(test.arr)-1 {
			isPeak = test.arr[result] > test.arr[result-1]
		} else {
			isPeak = test.arr[result] > test.arr[result-1] && test.arr[result] > test.arr[result+1]
		}

		if !isPeak {
			t.Errorf("FindPeakElement(%v) = %d; not a valid peak", test.arr, result)
		}
	}
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{16, 4},
		{1, 1},
		{0, 0},
		{2, 1},
		{15, 3},
	}

	for _, test := range tests {
		result := MySqrt(test.x)
		if result != test.expected {
			t.Errorf("MySqrt(%d) = %d; expected %d", test.x, result, test.expected)
		}
	}
}

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 3, 5, 6}

	tests := []struct {
		target   int
		expected int
	}{
		{5, 2}, // Found at index 2
		{2, 1}, // Insert at index 1
		{7, 4}, // Insert at end
		{0, 0}, // Insert at beginning
		{4, 2}, // Insert between 3 and 5
	}

	for _, test := range tests {
		result := SearchInsert(arr, test.target)
		if result != test.expected {
			t.Errorf("SearchInsert(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
		}
	}
}

// Benchmark tests
func BenchmarkBinarySearch(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i * 2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, 500)
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i * 2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchRecursive(arr, 500, 0, len(arr)-1)
	}
}
