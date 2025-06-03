# Merge Intervals Algorithm 📅

## Table of Contents
1. [Overview](#overview)
2. [Core Concept](#core-concept)
3. [Algorithm Variations](#algorithm-variations)
4. [Time & Space Complexity](#time--space-complexity)
5. [Implementation Patterns](#implementation-patterns)
6. [Common Problems](#common-problems)
7. [Step-by-Step Examples](#step-by-step-examples)
8. [Tips and Tricks](#tips-and-tricks)

---

## Overview

The **Merge Intervals Algorithm** is a fundamental technique used to handle overlapping intervals in a collection. It's widely used in scheduling problems, calendar applications, resource allocation, and range-based operations.

### Key Applications:
- 📅 **Calendar scheduling** - Merging overlapping appointments
- 🏢 **Meeting room allocation** - Optimizing room usage
- 📊 **Data processing** - Merging overlapping time ranges
- 🎯 **Range queries** - Combining overlapping ranges
- 🔄 **Resource management** - Handling conflicting reservations

---

## Core Concept

### What are Intervals?
An interval is represented as `[start, end]` where:
- `start` ≤ `end`
- Intervals can overlap, touch, or be completely separate

### Overlap Conditions:
Two intervals `[a, b]` and `[c, d]` overlap if:
- `a ≤ d` and `c ≤ b`

### Example:
```
Input:  [[1,3], [2,6], [8,10], [15,18]]
Output: [[1,6], [8,10], [15,18]]

Visualization:
1----3     →  1-------6
  2------6
            8---10  →  8---10
                  15----18  →  15----18
```

---

## Algorithm Variations

### 1. Basic Merge Intervals 🔄
**Problem**: Given a collection of intervals, merge all overlapping intervals.

**Approach**:
1. Sort intervals by start time
2. Iterate through sorted intervals
3. Merge overlapping intervals

```go
func MergeIntervals(intervals []Interval) []Interval {
    // Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    merged := []Interval{intervals[0]}
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start <= merged[len(merged)-1].End {
            // Merge intervals
            merged[len(merged)-1].End = max(merged[len(merged)-1].End, intervals[i].End)
        } else {
            // No overlap, add new interval
            merged = append(merged, intervals[i])
        }
    }
    
    return merged
}
```

### 2. Insert Interval ➕
**Problem**: Insert a new interval into a list of sorted non-overlapping intervals.

**Approach**:
1. Add all intervals that end before new interval starts
2. Merge all overlapping intervals with new interval
3. Add remaining intervals

```go
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
    result := []Interval{}
    i := 0
    
    // Add intervals that end before newInterval starts
    for i < len(intervals) && intervals[i].End < newInterval.Start {
        result = append(result, intervals[i])
        i++
    }
    
    // Merge overlapping intervals
    for i < len(intervals) && intervals[i].Start <= newInterval.End {
        newInterval.Start = min(newInterval.Start, intervals[i].Start)
        newInterval.End = max(newInterval.End, intervals[i].End)
        i++
    }
    result = append(result, newInterval)
    
    // Add remaining intervals
    for i < len(intervals) {
        result = append(result, intervals[i])
        i++
    }
    
    return result
}
```

### 3. Meeting Rooms I 🏢
**Problem**: Determine if a person can attend all meetings (no overlaps).

**Approach**:
1. Sort meetings by start time
2. Check for any overlaps between consecutive meetings

```go
func CanAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    
    for i := 1; i < len(intervals); i++ {
        if intervals[i].Start < intervals[i-1].End {
            return false // Overlap found
        }
    }
    
    return true
}
```

### 4. Meeting Rooms II 🏢🏢
**Problem**: Find minimum number of meeting rooms needed.

**Approach**: Use **Event Sorting** technique
1. Create separate arrays for start and end times
2. Sort both arrays
3. Use two pointers to track concurrent meetings

```go
func MinMeetingRooms(intervals []Interval) int {
    starts := make([]int, len(intervals))
    ends := make([]int, len(intervals))
    
    for i, interval := range intervals {
        starts[i] = interval.Start
        ends[i] = interval.End
    }
    
    sort.Ints(starts)
    sort.Ints(ends)
    
    rooms := 0
    maxRooms := 0
    startPtr, endPtr := 0, 0
    
    for startPtr < len(starts) {
        if starts[startPtr] < ends[endPtr] {
            rooms++ // Meeting starts
            startPtr++
        } else {
            rooms-- // Meeting ends
            endPtr++
        }
        maxRooms = max(maxRooms, rooms)
    }
    
    return maxRooms
}
```

### 5. Interval Intersection 🔀
**Problem**: Find intersection of two interval lists.

**Approach**: Use **Two Pointers**
1. Compare intervals from both lists
2. Find overlap using max(start1, start2) and min(end1, end2)
3. Move pointer of interval that ends first

```go
func IntervalIntersection(list1, list2 []Interval) []Interval {
    result := []Interval{}
    i, j := 0, 0
    
    for i < len(list1) && j < len(list2) {
        start := max(list1[i].Start, list2[j].Start)
        end := min(list1[i].End, list2[j].End)
        
        if start <= end {
            result = append(result, Interval{Start: start, End: end})
        }
        
        if list1[i].End < list2[j].End {
            i++
        } else {
            j++
        }
    }
    
    return result
}
```

### 6. Remove Covered Intervals 🗑️
**Problem**: Remove intervals that are completely covered by other intervals.

**Approach**: **Custom Sorting**
1. Sort by start time, then by end time (descending)
2. Keep intervals that extend beyond previous ones

```go
func RemoveCoveredIntervals(intervals []Interval) []Interval {
    // Sort by start, then by end (descending)
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i].Start == intervals[j].Start {
            return intervals[i].End > intervals[j].End
        }
        return intervals[i].Start < intervals[j].Start
    })
    
    result := []Interval{}
    
    for _, current := range intervals {
        if len(result) == 0 || current.End > result[len(result)-1].End {
            result = append(result, current)
        }
    }
    
    return result
}
```

---

## Time & Space Complexity

| Algorithm | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| Basic Merge | O(n log n) | O(1) extra | Sorting dominates |
| Insert Interval | O(n) | O(1) extra | Already sorted input |
| Meeting Rooms I | O(n log n) | O(1) extra | Sorting required |
| Meeting Rooms II | O(n log n) | O(n) | Separate arrays |
| Interval Intersection | O(m + n) | O(1) extra | Two pointers |
| Remove Covered | O(n log n) | O(1) extra | Custom sorting |

### Why O(n log n)?
- **Sorting step** is usually the bottleneck
- Most interval problems require sorted input
- Once sorted, the merge process is O(n)

---

## Implementation Patterns

### Pattern 1: Sort + Merge
```go
// 1. Sort intervals by start time
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i].Start < intervals[j].Start
})

// 2. Iterate and merge
for i := 1; i < len(intervals); i++ {
    if canMerge(intervals[i-1], intervals[i]) {
        // Merge logic
    }
}
```

### Pattern 2: Event Sorting
```go
// 1. Separate start and end events
events := []Event{}
for _, interval := range intervals {
    events = append(events, Event{Time: interval.Start, Type: START})
    events = append(events, Event{Time: interval.End, Type: END})
}

// 2. Sort events by time
sort.Slice(events, func(i, j int) bool {
    if events[i].Time == events[j].Time {
        return events[i].Type < events[j].Type // END before START
    }
    return events[i].Time < events[j].Time
})
```

### Pattern 3: Two Pointers
```go
// For intersection/comparison of two interval lists
i, j := 0, 0
for i < len(list1) && j < len(list2) {
    // Process current intervals
    // Move pointer of interval that ends first
    if list1[i].End < list2[j].End {
        i++
    } else {
        j++
    }
}
```

---

## Common Problems

### 1. **Merge Intervals** (LeetCode 56)
- Merge all overlapping intervals
- **Pattern**: Sort + Merge

### 2. **Insert Interval** (LeetCode 57)
- Insert and merge in sorted list
- **Pattern**: Three-phase processing

### 3. **Meeting Rooms** (LeetCode 252)
- Check if all meetings can be attended
- **Pattern**: Sort + Check overlap

### 4. **Meeting Rooms II** (LeetCode 253)
- Minimum meeting rooms needed
- **Pattern**: Event sorting or Priority Queue

### 5. **Interval List Intersections** (LeetCode 986)
- Find intersection of two lists
- **Pattern**: Two pointers

### 6. **Remove Covered Intervals** (LeetCode 1288)
- Remove intervals covered by others
- **Pattern**: Custom sorting

### 7. **Non-overlapping Intervals** (LeetCode 435)
- Minimum removals to make non-overlapping
- **Pattern**: Greedy + Sort by end time

---

## Step-by-Step Examples

### Example 1: Basic Merge
```
Input: [[1,3], [2,6], [8,10], [15,18]]

Step 1: Sort by start time (already sorted)
[[1,3], [2,6], [8,10], [15,18]]

Step 2: Initialize result with first interval
result = [[1,3]]

Step 3: Process [2,6]
- 2 ≤ 3? Yes → Merge
- result = [[1,6]]

Step 4: Process [8,10]
- 8 ≤ 6? No → Add new
- result = [[1,6], [8,10]]

Step 5: Process [15,18]
- 15 ≤ 10? No → Add new
- result = [[1,6], [8,10], [15,18]]
```

### Example 2: Meeting Rooms II
```
Input: [[0,30], [5,10], [15,20]]

Step 1: Create start and end arrays
starts = [0, 5, 15]
ends = [30, 10, 20]

Step 2: Sort both arrays
starts = [0, 5, 15]
ends = [10, 20, 30]

Step 3: Process events
Time 0: Meeting starts → rooms = 1
Time 5: Meeting starts → rooms = 2
Time 10: Meeting ends → rooms = 1
Time 15: Meeting starts → rooms = 2
Time 20: Meeting ends → rooms = 1
Time 30: Meeting ends → rooms = 0

Maximum rooms needed: 2
```

---

## Tips and Tricks

### 🎯 **Sorting Strategy**
- **By start time**: Most common approach
- **By end time**: For greedy algorithms (minimum removals)
- **Custom sorting**: When dealing with covered intervals

### 🎯 **Overlap Detection**
```go
// Two intervals [a,b] and [c,d] overlap if:
func isOverlap(a, b, c, d int) bool {
    return a <= d && c <= b
}

// Simplified for sorted intervals:
func isOverlap(prev, curr Interval) bool {
    return curr.Start <= prev.End
}
```

### 🎯 **Merge Logic**
```go
// When merging intervals [a,b] and [c,d]:
newStart := min(a, c)  // Usually just 'a' if sorted
newEnd := max(b, d)
```

### 🎯 **Edge Cases to Consider**
- Empty input
- Single interval
- All intervals overlap
- No intervals overlap
- Adjacent intervals (touching but not overlapping)

### 🎯 **Common Mistakes**
- Forgetting to sort input
- Wrong overlap condition
- Not handling edge cases
- Incorrect merge logic for end times

### 🎯 **Optimization Techniques**
- **In-place merging**: Modify input array instead of creating new one
- **Early termination**: Stop when no more overlaps possible
- **Binary search**: For insertion in sorted lists

---

## Practice Problems

### Beginner Level 🟢
1. **Merge Intervals** - Basic merging
2. **Meeting Rooms** - Simple overlap detection
3. **Insert Interval** - Single insertion

### Intermediate Level 🟡
1. **Meeting Rooms II** - Resource allocation
2. **Interval List Intersections** - Two-list processing
3. **Remove Covered Intervals** - Advanced sorting

### Advanced Level 🔴
1. **Non-overlapping Intervals** - Greedy optimization
2. **Employee Free Time** - Complex interval operations
3. **Range Module** - Dynamic interval management

---

## Summary

The **Merge Intervals Algorithm** is essential for handling time-based and range-based problems. The key insights are:

1. **Sorting** is usually the first step
2. **Greedy approach** works for most interval problems
3. **Two pointers** are useful for interval intersections
4. **Event sorting** helps with complex scheduling problems

Master these patterns, and you'll be able to solve most interval-related problems efficiently! 🚀 