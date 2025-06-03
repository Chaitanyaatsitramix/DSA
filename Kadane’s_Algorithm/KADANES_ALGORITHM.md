# Kadane's Algorithm

**Kadane's Algorithm** is a dynamic programming algorithm used to solve the **Maximum Subarray Problem** - finding the contiguous subarray with the largest sum in an array of numbers. Proposed by Jay Kadane in 1984, it runs in O(n) time complexity and O(1) space complexity.

## The Problem

Given an array of integers (both positive and negative), find the contiguous subarray that has the maximum sum.

**Example**: 
- Array: `[-2, -3, 4, -1, -2, 1, 5, -3]`
- Maximum subarray: `[4, -1, -2, 1, 5]` with sum = `7`

## Core Principle

At each position, we have two choices:
1. **Start a new subarray** from the current element
2. **Extend the existing subarray** by including the current element

**Key Insight**: If the sum accumulated so far becomes negative, it's better to start fresh from the current element rather than carrying the negative sum forward.

## Algorithm Logic

```
maxSoFar = arr[0]          // Maximum sum found so far
maxEndingHere = arr[0]     // Maximum sum ending at current position

for i = 1 to n-1:
    maxEndingHere = max(arr[i], maxEndingHere + arr[i])
    maxSoFar = max(maxSoFar, maxEndingHere)
```

## Methods of Kadane's Algorithm

### 1. Classic Kadane's Algorithm 🏆

**Purpose**: Find the maximum sum of contiguous subarray

```go
func MaxSubarraySum(arr []int) int {
    maxSoFar := arr[0]
    maxEndingHere := arr[0]
    
    for i := 1; i < len(arr); i++ {
        maxEndingHere = max(arr[i], maxEndingHere + arr[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    
    return maxSoFar
}
```

**Time Complexity**: O(n)  
**Space Complexity**: O(1)

### 2. Kadane's with Indices 📍

**Purpose**: Find maximum sum AND track the start and end positions

```go
func MaxSubarrayWithIndices(arr []int) (int, int, int) {
    maxSoFar := arr[0]
    maxEndingHere := arr[0]
    start, end, tempStart := 0, 0, 0
    
    for i := 1; i < len(arr); i++ {
        if arr[i] > maxEndingHere + arr[i] {
            maxEndingHere = arr[i]
            tempStart = i
        } else {
            maxEndingHere = maxEndingHere + arr[i]
        }
        
        if maxEndingHere > maxSoFar {
            maxSoFar = maxEndingHere
            start = tempStart
            end = i
        }
    }
    
    return maxSoFar, start, end
}
```

### 3. Circular Kadane's Algorithm 🔄

**Purpose**: Handle circular arrays where the last element connects to the first

**Strategy**: 
- **Case 1**: Maximum subarray is non-circular (normal Kadane's)
- **Case 2**: Maximum subarray is circular (total sum - minimum subarray)

```go
func MaxCircularSubarray(arr []int) int {
    // Case 1: Normal Kadane's
    maxKadane := MaxSubarraySum(arr)
    
    // Case 2: Circular - find minimum subarray and subtract from total
    arraySum := 0
    for i := 0; i < len(arr); i++ {
        arraySum += arr[i]
        arr[i] = -arr[i]  // Negate for finding minimum
    }
    
    minSubarraySum := -MaxSubarraySum(arr)  // Find min of original
    maxCircular := arraySum - minSubarraySum
    
    // Restore array and handle edge cases
    for i := 0; i < len(arr); i++ {
        arr[i] = -arr[i]
    }
    
    if maxCircular == 0 {
        return maxKadane
    }
    
    return max(maxKadane, maxCircular)
}
```

## Variations of Kadane's Algorithm

### 1. Maximum Product Subarray

**Challenge**: Handle negative numbers that can make products larger

```go
func MaxProductSubarray(arr []int) int {
    maxSoFar := arr[0]
    maxEndingHere := arr[0]
    minEndingHere := arr[0]  // Track minimum for negative numbers
    
    for i := 1; i < len(arr); i++ {
        temp := maxEndingHere
        maxEndingHere = max(arr[i], max(arr[i]*maxEndingHere, arr[i]*minEndingHere))
        minEndingHere = min(arr[i], min(arr[i]*temp, arr[i]*minEndingHere))
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    
    return maxSoFar
}
```

### 2. Maximum Sum with Constraints

**Example**: No three consecutive elements

```go
func MaxSumNoThreeConsecutive(arr []int) int {
    n := len(arr)
    if n <= 2 {
        return sum(arr)
    }
    
    dp := make([]int, n)
    dp[0] = arr[0]
    dp[1] = arr[0] + arr[1]
    dp[2] = max(dp[1], max(arr[1]+arr[2], arr[0]+arr[2]))
    
    for i := 3; i < n; i++ {
        dp[i] = max(dp[i-1], max(dp[i-2]+arr[i], dp[i-3]+arr[i-1]+arr[i]))
    }
    
    return dp[n-1]
}
```

## Step-by-Step Example

**Array**: `[-2, -3, 4, -1, -2, 1, 5, -3]`

| i | arr[i] | maxEndingHere | maxSoFar | Explanation |
|---|--------|---------------|----------|-------------|
| 0 | -2     | -2            | -2       | Initialize |
| 1 | -3     | -3            | -2       | max(-3, -2-3) = -3 |
| 2 | 4      | 4             | 4        | max(4, -3+4) = 4 (start new) |
| 3 | -1     | 3             | 4        | max(-1, 4-1) = 3 (extend) |
| 4 | -2     | 1             | 4        | max(-2, 3-2) = 1 (extend) |
| 5 | 1      | 2             | 4        | max(1, 1+1) = 2 (extend) |
| 6 | 5      | 7             | 7        | max(5, 2+5) = 7 (extend) |
| 7 | -3     | 4             | 7        | max(-3, 7-3) = 4 (extend) |

**Result**: Maximum sum = 7, Subarray = `[4, -1, -2, 1, 5]`

## When to Use Kadane's Algorithm

✅ **Perfect for**:
- **Maximum/Minimum Subarray Sum**: Classic use case
- **Maximum Product Subarray**: With modifications for products
- **Optimization Problems**: Any problem involving contiguous sequences
- **Dynamic Programming**: When optimal substructure exists

❌ **Not suitable for**:
- **Non-contiguous subsequences**: Elements don't need to be adjacent
- **Multiple subarrays**: When you need k non-overlapping subarrays
- **Complex constraints**: When simple greedy choice doesn't work

## Key Insights

1. **Greedy Choice**: At each step, make the locally optimal choice
2. **Optimal Substructure**: Optimal solution contains optimal subsolutions
3. **Negative Reset**: When accumulated sum becomes negative, start fresh
4. **Single Pass**: Only need one traversal through the array

## Time & Space Complexity

- **Time Complexity**: O(n) - Single pass through the array
- **Space Complexity**: O(1) - Only using a few variables

## Common Mistakes

1. **Forgetting negative arrays**: When all elements are negative, return the maximum element
2. **Index tracking**: Be careful with start/end index updates
3. **Circular arrays**: Remember to handle the wrap-around case
4. **Product variations**: Track both maximum and minimum for negative numbers

## Extensions and Applications

1. **2D Kadane's**: Maximum sum rectangle in 2D matrix
2. **K-Kadane's**: Maximum sum of k non-overlapping subarrays
3. **Kadane's with queries**: Range maximum subarray queries
4. **Stock market**: Maximum profit from stock transactions

## Problem-Solving Template

```go
func kadaneTemplate(arr []int) int {
    // Initialize with first element
    maxSoFar := arr[0]
    maxEndingHere := arr[0]
    
    // Iterate through rest of array
    for i := 1; i < len(arr); i++ {
        // Decision: extend or start new
        maxEndingHere = max(arr[i], maxEndingHere + arr[i])
        
        // Update global maximum
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    
    return maxSoFar
}
```

Kadane's Algorithm is fundamental in dynamic programming and appears frequently in coding interviews. Master the basic pattern and its variations to solve a wide range of optimization problems efficiently! 