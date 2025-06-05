# Quickselect Algorithm

The Quickselect algorithm is an efficient selection algorithm to find the kth smallest element in an unordered list. It is related to the Quicksort sorting algorithm and uses similar partitioning techniques.

## Overview

Quickselect works by selecting a pivot element and partitioning the array around it, similar to Quicksort. However, instead of recursing on both sides like Quicksort, it only recurses on the side that contains the kth element.

### Time Complexity
- Average case: O(n)
- Worst case: O(n²)
- Best case: O(n)

### Space Complexity
- O(1) extra space (in-place algorithm)
- O(log n) stack space for recursion

## Implementation Methods

This implementation includes two main approaches:

1. **Basic Quickselect**
   - Uses Lomuto's partitioning scheme
   - Simple and efficient for most cases
   - Average case O(n) time complexity
   - May degrade to O(n²) in worst case

2. **Median of Medians**
   - Guaranteed O(n) worst-case time complexity
   - More complex implementation
   - Uses deterministic pivot selection
   - Better for critical applications requiring guaranteed performance

## Usage Examples

```go
// Finding the kth smallest element
arr := []int{7, 10, 4, 3, 20, 15}
k := 3 // Find the 4th smallest element (0-based index)
result := quickselect.QuickSelect(arr, k)

// Using Median of Medians for guaranteed performance
result = quickselect.MedianOfMedians(arr, k)
```

## Common Applications

1. **Finding Medians**: Efficiently find the median of an unsorted array
2. **Order Statistics**: Find any kth order statistic in an array
3. **Partial Sorting**: When only a specific position needs to be correct
4. **Top-K Problems**: Finding the kth largest/smallest elements

## Implementation Details

### Basic Quickselect
1. Choose a pivot element
2. Partition array around pivot
3. Compare k with pivot position:
   - If equal: return pivot element
   - If less: recurse on left subarray
   - If greater: recurse on right subarray

### Median of Medians
1. Divide array into groups of 5
2. Find median of each group
3. Recursively find median of medians
4. Use this as pivot for partitioning
5. Recurse on appropriate side

## Edge Cases Handled

- Empty arrays
- Single element arrays
- Arrays with duplicate elements
- Invalid k values
- Arrays requiring worst-case performance

## Performance Considerations

1. **Pivot Selection**
   - Random selection for basic Quickselect
   - Median of medians for guaranteed performance

2. **Partitioning Strategy**
   - Lomuto's partition scheme
   - Handles duplicate elements correctly

3. **Recursion Depth**
   - Limited to O(log n) on average
   - Can be O(n) in worst case for basic Quickselect 