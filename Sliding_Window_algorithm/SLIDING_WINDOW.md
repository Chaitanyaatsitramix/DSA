# Sliding Window Algorithm

The **Sliding Window Algorithm** is a computational technique used to efficiently solve problems involving contiguous sequences (subarrays/substrings) in arrays or strings. It maintains a "window" that slides across the data structure, reducing time complexity from O(n²) or O(n³) to O(n).

## How It Works

Instead of recalculating everything for each subarray, the sliding window:
1. **Maintains a window** of elements
2. **Slides the window** by removing elements from one end and adding to the other
3. **Updates the result** incrementally

## Two Main Methods

### 1. Fixed Size Window 🪟

**Description**: The window size remains constant throughout the algorithm.

**Pattern**:
```go
windowSum := 0
// Calculate first window
for i := 0; i < k; i++ {
    windowSum += arr[i]
}

// Slide the window
for i := k; i < len(arr); i++ {
    windowSum = windowSum - arr[i-k] + arr[i]  // Remove left, add right
    // Process current window
}
```

**Use Cases**:
- Maximum/minimum sum of subarray of size k
- Average of all subarrays of size k
- Maximum elements in sliding window
- Finding anagrams in a string

**Example - Maximum Sum Subarray**:
```go
func MaxSumFixedWindow(arr []int, k int) int {
    windowSum := 0
    // Calculate sum of first window
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
```

### 2. Variable Size Window 📏

**Description**: Window size can expand and contract based on certain conditions.

**Pattern**:
```go
left := 0
for right := 0; right < len(arr); right++ {
    // Expand window by including arr[right]
    
    while window_condition_violated {
        // Shrink window from left
        left++
    }
    
    // Update result with current valid window
}
```

**Sub-types**:

#### **Expanding Window**
- Grows until a condition is violated
- Used for finding longest valid sequences

#### **Shrinking Window**  
- Contracts to maintain validity
- Used for finding minimum valid sequences

**Use Cases**:
- Longest substring with k distinct characters
- Minimum window substring
- Longest substring without repeating characters
- Subarray with given sum
- Maximum consecutive 1s

**Example - Longest Substring K Distinct**:
```go
func LongestSubstringKDistinct(s string, k int) int {
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
```

## Key Advantages

1. **Time Efficiency**: Reduces time complexity from O(n²) to O(n)
2. **Space Efficiency**: Usually O(1) or O(k) space complexity
3. **Intuitive**: Easy to understand and visualize
4. **Versatile**: Works for arrays, strings, and linked lists

## When to Use Sliding Window

✅ **Perfect for**:
- **Contiguous subarrays/substrings**: Problems involving consecutive elements
- **Optimization problems**: Finding maximum/minimum of subarrays
- **String matching**: Pattern matching and substring problems
- **Sum/average calculations**: Window-based calculations

❌ **Not suitable for**:
- **Non-contiguous subsequences**: Elements don't need to be adjacent
- **Sorting required**: When array order needs to be changed
- **Global optimization**: When local optimization isn't sufficient

## Common Problem Patterns

### Pattern 1: Fixed Window Size
```go
// Template for fixed size window
windowSum := 0
for i := 0; i < k; i++ {
    windowSum += arr[i]
}
result := windowSum

for i := k; i < len(arr); i++ {
    windowSum = windowSum - arr[i-k] + arr[i]
    result = updateResult(result, windowSum)
}
```

### Pattern 2: Variable Window - Find Maximum
```go
// Template for longest valid window
left := 0
maxLen := 0

for right := 0; right < len(arr); right++ {
    // Add arr[right] to window
    
    while window_invalid {
        // Remove arr[left] from window
        left++
    }
    
    maxLen = max(maxLen, right-left+1)
}
```

### Pattern 3: Variable Window - Find Minimum
```go
// Template for shortest valid window
left := 0
minLen := math.MaxInt32

for right := 0; right < len(arr); right++ {
    // Add arr[right] to window
    
    while window_valid {
        minLen = min(minLen, right-left+1)
        // Remove arr[left] from window
        left++
    }
}
```

## Problem-Solving Strategy

1. **Identify the problem type**:
   - Fixed size window vs Variable size window
   - Maximum vs Minimum optimization

2. **Define window validity**:
   - What makes a window valid/invalid?
   - What condition triggers expansion/contraction?

3. **Choose the right template**:
   - Use fixed size for known window size
   - Use variable size for dynamic window requirements

4. **Implement window operations**:
   - How to add elements to window?
   - How to remove elements from window?
   - How to check window validity?

5. **Track the result**:
   - When to update the result?
   - What to return (length, sum, substring, etc.)?

## Time & Space Complexity

**Time Complexity**: O(n)
- Each element is visited at most twice (once by right pointer, once by left pointer)

**Space Complexity**: 
- O(1) for simple problems
- O(k) where k is the size of auxiliary data structures (like hashmaps)

## Tips for Implementation

1. **Use appropriate data structures**:
   - HashMap for character/element counting
   - Deque for maintaining window extremes
   - Simple variables for sum/count tracking

2. **Handle edge cases**:
   - Empty arrays/strings
   - Window size larger than array
   - All elements satisfy/violate condition

3. **Optimize space**:
   - Use arrays instead of hashmaps when possible
   - Clean up data structures when elements are removed

4. **Debug effectively**:
   - Print window boundaries and contents
   - Verify window validity at each step
   - Check result updates

The sliding window technique is essential for solving array and string problems efficiently. Master both fixed and variable size windows to tackle a wide variety of algorithmic challenges! 