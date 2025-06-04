# Binary Search Algorithm 🔍

## Table of Contents
1. [Overview](#overview)
2. [Core Concept](#core-concept)
3. [Algorithm Variations](#algorithm-variations)
4. [Time & Space Complexity](#time--space-complexity)
5. [Implementation Methods](#implementation-methods)
6. [Common Problems & Patterns](#common-problems--patterns)
7. [Step-by-Step Examples](#step-by-step-examples)
8. [Tips and Tricks](#tips-and-tricks)

---

## Overview

**Binary Search** is a highly efficient searching algorithm that works on sorted arrays. It uses the "divide and conquer" strategy to repeatedly divide the search space in half, eliminating half of the remaining elements with each comparison.

### Key Characteristics:
- 🎯 **Prerequisite**: Array must be sorted
- ⚡ **Time Complexity**: O(log n)
- 💾 **Space Complexity**: O(1) iterative, O(log n) recursive
- 🔄 **Strategy**: Divide and conquer

### Applications:
- 🔍 **Searching in sorted arrays**
- 📊 **Finding insertion points**
- 🎯 **Range queries**
- 🔢 **Mathematical computations** (square root, power)
- 📈 **Optimization problems**

---

## Core Concept

### How Binary Search Works:

1. **Compare** the target with the middle element
2. **Eliminate** half of the search space based on comparison
3. **Repeat** until target is found or search space is empty

### Visual Example:
```
Array: [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
Target: 7

Step 1: [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
        left=0, right=9, mid=4
        arr[4] = 9 > 7, so search left half

Step 2: [1, 3, 5, 7, 9]
        left=0, right=3, mid=1
        arr[1] = 3 < 7, so search right half

Step 3: [5, 7]
        left=2, right=3, mid=2
        arr[2] = 5 < 7, so search right half

Step 4: [7]
        left=3, right=3, mid=3
        arr[3] = 7 = 7, FOUND!
```

---

## Algorithm Variations

### 1. Standard Binary Search 🎯
**Problem**: Find the exact position of target in sorted array.

**Returns**: Index of target or -1 if not found.

```go
func BinarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2  // Prevents overflow
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1  // Target not found
}
```

### 2. Lower Bound (First Occurrence) ⬇️
**Problem**: Find the first occurrence of target or the smallest index where we can insert target.

**Returns**: Leftmost position where target can be inserted.

```go
func LowerBound(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}
```

### 3. Upper Bound (After Last Occurrence) ⬆️
**Problem**: Find the position after the last occurrence of target.

**Returns**: Position where we can insert target to maintain sorted order.

```go
func UpperBound(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}
```

### 4. First and Last Position 🎯🎯
**Problem**: Find the range [first, last] of target occurrences.

```go
func SearchRange(arr []int, target int) [2]int {
    first := findFirst(arr, target)
    if first == -1 {
        return [2]int{-1, -1}
    }
    
    last := findLast(arr, target)
    return [2]int{first, last}
}

func findFirst(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            result = mid
            right = mid - 1  // Continue searching left
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}

func findLast(arr []int, target int) int {
    left, right := 0, len(arr)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            result = mid
            left = mid + 1  // Continue searching right
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}
```

### 5. Peak Element 🏔️
**Problem**: Find any peak element (element greater than its neighbors).

```go
func FindPeakElement(arr []int) int {
    left, right := 0, len(arr)-1
    
    for left < right {
        mid := left + (right-left)/2
        
        if arr[mid] > arr[mid+1] {
            right = mid  // Peak is on the left side or at mid
        } else {
            left = mid + 1  // Peak is on the right side
        }
    }
    
    return left
}
```

### 6. Rotated Sorted Array 🔄
**Problem**: Search in rotated sorted array.

```go
func SearchRotated(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        }
        
        // Check which part is sorted
        if arr[left] <= arr[mid] {  // Left part is sorted
            if arr[left] <= target && target < arr[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {  // Right part is sorted
            if arr[mid] < target && target <= arr[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}
```

---

## Time & Space Complexity

| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Standard Search | O(log n) | O(1) iterative | Most efficient search |
| Recursive Search | O(log n) | O(log n) | Call stack overhead |
| Lower/Upper Bound | O(log n) | O(1) | Same as standard |
| Range Search | O(log n) | O(1) | Two binary searches |
| Peak Finding | O(log n) | O(1) | Modified binary search |
| Rotated Array | O(log n) | O(1) | With extra conditions |

### Why O(log n)?
- Each iteration eliminates **half** of the remaining elements
- Maximum iterations = log₂(n)
- Example: Array of 1000 elements → max 10 iterations

---

## Implementation Methods

### Method 1: Iterative Approach (Recommended) 🔄
```go
func BinarySearchIterative(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}
```

**Advantages:**
- ✅ O(1) space complexity
- ✅ No function call overhead
- ✅ Better for large datasets

### Method 2: Recursive Approach 📞
```go
func BinarySearchRecursive(arr []int, target, left, right int) int {
    if left > right {
        return -1
    }
    
    mid := left + (right-left)/2
    
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        return BinarySearchRecursive(arr, target, mid+1, right)
    } else {
        return BinarySearchRecursive(arr, target, left, mid-1)
    }
}

// Wrapper function
func BinarySearch(arr []int, target int) int {
    return BinarySearchRecursive(arr, target, 0, len(arr)-1)
}
```

**Advantages:**
- ✅ More intuitive and readable
- ✅ Easier to understand the divide-and-conquer concept

**Disadvantages:**
- ❌ O(log n) space due to call stack
- ❌ Function call overhead

### Method 3: Template Method (Generic) 🎯
```go
func BinarySearchTemplate(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if check(arr, mid, target) == 0 {
            return mid
        } else if check(arr, mid, target) < 0 {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return -1
}

func check(arr []int, mid, target int) int {
    if arr[mid] == target {
        return 0
    } else if arr[mid] < target {
        return -1
    } else {
        return 1
    }
}
```

---

## Common Problems & Patterns

### Pattern 1: Find Exact Match 🎯
**Problems**: 
- Search in sorted array
- Search in 2D matrix

### Pattern 2: Find Boundaries 📍
**Problems**: 
- First/Last occurrence
- Search range
- Insert position

### Pattern 3: Find Peak/Valley 🏔️
**Problems**: 
- Peak element
- Valley element
- Local maximum/minimum

### Pattern 4: Rotated Arrays 🔄
**Problems**: 
- Search in rotated sorted array
- Find minimum in rotated array
- Find rotation point

### Pattern 5: Mathematical Applications 🔢
**Problems**: 
- Square root
- Power function
- Find kth smallest/largest

### Pattern 6: Condition-based Search ✅
**Problems**: 
- Capacity to ship packages
- Smallest divisor
- Minimum eating speed

---

## Step-by-Step Examples

### Example 1: Standard Binary Search
```
Array: [2, 5, 8, 12, 16, 23, 38, 45, 56, 67, 78]
Target: 23

Iteration 1:
left=0, right=10, mid=5
arr[5] = 23 = target ✅
FOUND at index 5!
```

### Example 2: Target Not Found
```
Array: [1, 3, 5, 7, 9]
Target: 6

Iteration 1:
left=0, right=4, mid=2
arr[2] = 5 < 6, search right half
left = 3

Iteration 2:
left=3, right=4, mid=3
arr[3] = 7 > 6, search left half
right = 2

Now left > right (3 > 2), stop searching
Return -1 (not found)
```

### Example 3: First and Last Occurrence
```
Array: [1, 2, 2, 2, 3, 4, 5]
Target: 2

Finding First Occurrence:
- Keep searching left when target is found
- Result: index 1

Finding Last Occurrence:
- Keep searching right when target is found  
- Result: index 3

Range: [1, 3]
```

---

## Tips and Tricks

### 🎯 **Avoiding Integer Overflow**
```go
// ❌ Wrong way (can overflow)
mid := (left + right) / 2

// ✅ Correct way
mid := left + (right - left) / 2
```

### 🎯 **Boundary Conditions**
- Always check `left <= right` for exact search
- Use `left < right` for boundary finding
- Handle empty arrays

### 🎯 **Loop Invariants**
- **Standard Search**: `left <= right`
- **Lower Bound**: `left < right`
- **Upper Bound**: `left < right`

### 🎯 **Common Mistakes**
- Forgetting to update `left` or `right`
- Using wrong comparison operators
- Infinite loops due to incorrect boundary updates
- Not handling edge cases (empty array, single element)

### 🎯 **When to Use Which Variation**
- **Standard**: When you need exact match
- **Lower Bound**: When finding insertion point or first occurrence
- **Upper Bound**: When finding position after last occurrence
- **Modified**: When searching in rotated arrays or finding peaks

### 🎯 **Debug Techniques**
```go
func BinarySearchDebug(arr []int, target int) int {
    left, right := 0, len(arr)-1
    iteration := 1
    
    for left <= right {
        mid := left + (right-left)/2
        fmt.Printf("Iteration %d: left=%d, right=%d, mid=%d, arr[mid]=%d\n", 
                   iteration, left, right, mid, arr[mid])
        
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
        iteration++
    }
    
    return -1
}
```

---

## Practice Problems

### Beginner Level 🟢
1. **Binary Search** - Standard implementation
2. **Search Insert Position** - Lower bound
3. **Find First and Last Position** - Range search

### Intermediate Level 🟡
1. **Search in Rotated Sorted Array** - Modified binary search
2. **Find Peak Element** - Peak finding
3. **Search a 2D Matrix** - Extended binary search

### Advanced Level 🔴
1. **Median of Two Sorted Arrays** - Complex binary search
2. **Split Array Largest Sum** - Condition-based search
3. **Capacity to Ship Packages** - Optimization problems

---

## Summary

**Binary Search** is one of the most important algorithms to master:

### Key Takeaways:
1. **Prerequisite**: Array must be sorted
2. **Core Idea**: Eliminate half of search space each iteration
3. **Time Complexity**: O(log n) - extremely efficient
4. **Space Complexity**: O(1) for iterative, O(log n) for recursive
5. **Variations**: Standard, Lower/Upper bound, Peak finding, Rotated arrays

### Master These Patterns:
- **Exact Search**: Find specific element
- **Boundary Search**: Find insertion points or ranges
- **Peak Finding**: Find local maximum/minimum
- **Condition-based**: Solve optimization problems

Binary Search is not just about finding elements - it's a powerful technique for solving optimization problems and is the foundation for many advanced algorithms! 🚀 