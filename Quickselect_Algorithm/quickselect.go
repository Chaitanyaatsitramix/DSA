package quickselect

// QuickSelect finds the kth smallest element in an unordered array
// k is 0-based index (0 means smallest element, 1 means second smallest, etc.)
func QuickSelect(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}
	return quickSelectHelper(arr, 0, len(arr)-1, k)
}

// quickSelectHelper is the recursive helper function for QuickSelect
func quickSelectHelper(arr []int, left, right, k int) int {
	// If the array contains only one element, return that element
	if left == right {
		return arr[left]
	}

	// Get the pivot position
	pivotIdx := partition(arr, left, right)

	// If pivot is the kth element, return it
	if k == pivotIdx {
		return arr[k]
	}
	// If k is less than pivot index, search in left subarray
	if k < pivotIdx {
		return quickSelectHelper(arr, left, pivotIdx-1, k)
	}
	// If k is greater than pivot index, search in right subarray
	return quickSelectHelper(arr, pivotIdx+1, right, k)
}

// partition function using Lomuto's partition scheme
func partition(arr []int, left, right int) int {
	// Choose rightmost element as pivot
	pivot := arr[right]
	i := left - 1

	// Move elements smaller than pivot to the left side
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place pivot in its final position
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

// MedianOfMedians is an optimized version of QuickSelect that guarantees O(n) time complexity
func MedianOfMedians(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}
	return medianOfMediansHelper(arr, 0, len(arr)-1, k)
}

func medianOfMediansHelper(arr []int, left, right, k int) int {
	// If array is small enough, use regular quickselect
	if right-left+1 <= 5 {
		return quickSelectHelper(arr, left, right, k)
	}

	// Divide array into groups of 5 and find medians
	numGroups := (right - left + 1) / 5
	medians := make([]int, numGroups)

	for i := 0; i < numGroups; i++ {
		groupStart := left + i*5
		groupEnd := groupStart + 4
		if groupEnd > right {
			groupEnd = right
		}

		// Find median of current group
		medians[i] = findMedian(arr[groupStart : groupEnd+1])
	}

	// Find median of medians
	medianOfMedians := MedianOfMedians(medians, len(medians)/2)

	// Use median of medians as pivot
	pivotIdx := partitionAroundValue(arr, left, right, medianOfMedians)

	if k == pivotIdx {
		return arr[k]
	} else if k < pivotIdx {
		return medianOfMediansHelper(arr, left, pivotIdx-1, k)
	}
	return medianOfMediansHelper(arr, pivotIdx+1, right, k)
}

// findMedian finds the median of a small array (size <= 5)
func findMedian(arr []int) int {
	n := len(arr)
	// Sort small array (insertion sort is efficient for small arrays)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr[n/2]
}

// partitionAroundValue partitions array around a specific value
func partitionAroundValue(arr []int, left, right, value int) int {
	for i := left; i <= right; i++ {
		if arr[i] == value {
			arr[i], arr[right] = arr[right], arr[i]
			break
		}
	}
	return partition(arr, left, right)
}
