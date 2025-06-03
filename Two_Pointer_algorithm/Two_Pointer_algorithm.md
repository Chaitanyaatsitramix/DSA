# Two-Pointer Algorithm

The **Two-Pointer Algorithm** is a powerful technique used to solve array and string problems more efficiently. Instead of using nested loops (O(n²)), it uses two pointers to traverse the data structure, often achieving O(n) time complexity.

## Two Main Methods

### 1. Opposite Direction (Convergent) Pointers 🔄
- **Description**: Two pointers start at opposite ends of the array and move towards each other
- **Pattern**: `left = 0, right = n-1` → move inward until `left >= right`
- **Time Complexity**: O(n)
- **Space Complexity**: O(1)

**Common Use Cases:**
- Finding pairs with target sum (in sorted array)
- Palindrome checking
- Reversing arrays/strings
- Container with most water
- 3Sum problems

**Example - Two Sum in Sorted Array:**
```go
func twoSum(nums []int, target int) []int {
    left, right := 0, len(nums)-1
    
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return []int{left, right}
        } else if sum < target {
            left++  // Need bigger sum
        } else {
            right-- // Need smaller sum
        }
    }
    return []int{-1, -1} // Not found
}
```

### 2. Same Direction (Sliding Window) Pointers 📊
- **Description**: Both pointers start at the beginning and move in the same direction at different speeds
- **Pattern**: `slow = 0, fast = 0` → fast moves ahead, slow follows conditionally
- **Time Complexity**: O(n)
- **Space Complexity**: O(1)

**Common Use Cases:**
- Removing duplicates from sorted arrays
- Finding subarrays with specific properties
- Longest substring problems
- Maximum/minimum sliding window
- Fast and slow pointer technique (cycle detection)

**Example - Remove Duplicates:**
```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0 // Points to position of unique elements
    
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }
    
    return slow + 1 // Length of unique elements
}
```

## Key Advantages

1. **Efficiency**: Reduces time complexity from O(n²) to O(n)
2. **Space Optimization**: Uses O(1) extra space
3. **Simplicity**: Easy to understand and implement
4. **Versatility**: Works for many different problem types

## When to Use Two Pointers

- **Sorted Arrays**: When the array is sorted or can be sorted
- **Target Sum Problems**: Finding pairs/triplets that sum to a target
- **Palindromes**: Checking if strings/arrays are palindromes
- **Subarray Problems**: Finding subarrays with specific properties
- **Cycle Detection**: Using fast/slow pointers in linked lists

## Tips for Implementation

1. **Choose the Right Method**: 
   - Use opposite direction for sorted arrays and target sum problems
   - Use same direction for sliding window and duplicate removal

2. **Handle Edge Cases**:
   - Empty arrays
   - Single element arrays
   - All elements are the same

3. **Watch for Infinite Loops**:
   - Ensure pointers move in every iteration
   - Have proper exit conditions

4. **Consider Sorting**:
   - Sometimes sorting the array first enables two-pointer technique

## Common Patterns

### Pattern 1: Two Sum (Sorted Array)
```go
left, right := 0, len(arr)-1
for left < right {
    if condition_met {
        return result
    } else if need_larger_value {
        left++
    } else {
        right--
    }
}
```

### Pattern 2: Sliding Window
```go
left := 0
for right := 0; right < len(arr); right++ {
    // Expand window by including arr[right]
    
    for window_invalid {
        // Shrink window from left
        left++
    }
    
    // Update result with current window
}
```

### Pattern 3: Fast and Slow Pointers
```go
slow, fast := 0, 0
for fast < len(arr) {
    if condition {
        arr[slow] = arr[fast]
        slow++
    }
    fast++
}
```

The two-pointer technique is fundamental in competitive programming and technical interviews. Master both methods to solve a wide variety of array and string problems efficiently! 