package main

import (
	"fmt"
	"sort"
)

// Interval represents a time interval with start and end
type Interval struct {
	Start int
	End   int
}

// ========== BASIC MERGE INTERVALS ==========

// Example 1: Basic Merge Overlapping Intervals
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := &merged[len(merged)-1]

		// If current interval overlaps with the last merged interval
		if current.Start <= lastMerged.End {
			// Merge intervals by extending the end time
			if current.End > lastMerged.End {
				lastMerged.End = current.End
			}
		} else {
			// No overlap, add current interval to result
			merged = append(merged, current)
		}
	}

	return merged
}

// ========== INSERT INTERVAL ==========

// Example 2: Insert a new interval and merge overlapping ones
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
	result := []Interval{}
	i := 0
	n := len(intervals)

	// Add all intervals that end before newInterval starts
	for i < n && intervals[i].End < newInterval.Start {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals with newInterval
	for i < n && intervals[i].Start <= newInterval.End {
		// Merge with newInterval
		if intervals[i].Start < newInterval.Start {
			newInterval.Start = intervals[i].Start
		}
		if intervals[i].End > newInterval.End {
			newInterval.End = intervals[i].End
		}
		i++
	}

	// Add the merged interval
	result = append(result, newInterval)

	// Add remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// ========== MEETING ROOMS PROBLEMS ==========

// Example 3: Can attend all meetings (no overlapping intervals)
func CanAttendMeetings(intervals []Interval) bool {
	if len(intervals) <= 1 {
		return true
	}

	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// Check for overlaps
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false // Overlap found
		}
	}

	return true
}

// Example 4: Minimum number of meeting rooms needed
func MinMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	// Create separate arrays for start and end times
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval.Start
		ends[i] = interval.End
	}

	// Sort both arrays
	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0
	maxRooms := 0
	startPtr := 0
	endPtr := 0

	// Use two pointers to track concurrent meetings
	for startPtr < len(starts) {
		if starts[startPtr] < ends[endPtr] {
			// Meeting starts, need a room
			rooms++
			startPtr++
		} else {
			// Meeting ends, free a room
			rooms--
			endPtr++
		}

		// Track maximum rooms needed
		if rooms > maxRooms {
			maxRooms = rooms
		}
	}

	return maxRooms
}

// ========== INTERVAL OPERATIONS ==========

// Example 5: Find intersection of two interval lists
func IntervalIntersection(list1, list2 []Interval) []Interval {
	result := []Interval{}
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		// Find intersection
		start := max(list1[i].Start, list2[j].Start)
		end := min(list1[i].End, list2[j].End)

		// If there's a valid intersection
		if start <= end {
			result = append(result, Interval{Start: start, End: end})
		}

		// Move pointer of interval that ends first
		if list1[i].End < list2[j].End {
			i++
		} else {
			j++
		}
	}

	return result
}

// Example 6: Remove intervals that are covered by others
func RemoveCoveredIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort by start time, then by end time in descending order
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End > intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})

	result := []Interval{}

	for _, current := range intervals {
		// Check if current interval is covered by the last added interval
		if len(result) == 0 || current.End > result[len(result)-1].End {
			result = append(result, current)
		}
		// If current.End <= result[last].End, then current is covered
	}

	return result
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ========== DEMO FUNCTIONS ==========

// Demo function for Basic Merge Intervals
func DemoBasicMerge() {
	fmt.Println("=== BASIC MERGE INTERVALS ===")

	intervals := []Interval{
		{Start: 1, End: 3},
		{Start: 2, End: 6},
		{Start: 8, End: 10},
		{Start: 15, End: 18},
	}

	fmt.Printf("Input intervals: %v\n", intervals)
	merged := MergeIntervals(intervals)
	fmt.Printf("Merged intervals: %v\n\n", merged)
}

// Demo function for Insert Interval
func DemoInsertInterval() {
	fmt.Println("=== INSERT INTERVAL ===")

	intervals := []Interval{
		{Start: 1, End: 3},
		{Start: 6, End: 9},
	}
	newInterval := Interval{Start: 2, End: 5}

	fmt.Printf("Existing intervals: %v\n", intervals)
	fmt.Printf("New interval: %v\n", newInterval)
	result := InsertInterval(intervals, newInterval)
	fmt.Printf("After insertion: %v\n\n", result)
}

// Demo function for Meeting Rooms
func DemoMeetingRooms() {
	fmt.Println("=== MEETING ROOMS PROBLEMS ===")

	// Test can attend all meetings
	meetings1 := []Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}

	meetings2 := []Interval{
		{Start: 7, End: 10},
		{Start: 2, End: 4},
	}

	fmt.Printf("Meetings 1: %v\n", meetings1)
	fmt.Printf("Can attend all: %v\n", CanAttendMeetings(meetings1))

	fmt.Printf("Meetings 2: %v\n", meetings2)
	fmt.Printf("Can attend all: %v\n", CanAttendMeetings(meetings2))

	// Test minimum meeting rooms
	meetings3 := []Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}

	fmt.Printf("Meetings 3: %v\n", meetings3)
	fmt.Printf("Minimum rooms needed: %d\n\n", MinMeetingRooms(meetings3))
}

// Demo function for Interval Operations
func DemoIntervalOperations() {
	fmt.Println("=== INTERVAL OPERATIONS ===")

	// Test interval intersection
	list1 := []Interval{
		{Start: 0, End: 2},
		{Start: 5, End: 10},
		{Start: 13, End: 23},
		{Start: 24, End: 25},
	}

	list2 := []Interval{
		{Start: 1, End: 5},
		{Start: 8, End: 12},
		{Start: 15, End: 24},
		{Start: 25, End: 26},
	}

	fmt.Printf("List 1: %v\n", list1)
	fmt.Printf("List 2: %v\n", list2)
	intersection := IntervalIntersection(list1, list2)
	fmt.Printf("Intersection: %v\n", intersection)

	// Test remove covered intervals
	covered := []Interval{
		{Start: 1, End: 4},
		{Start: 3, End: 6},
		{Start: 2, End: 8},
	}

	fmt.Printf("Intervals with coverage: %v\n", covered)
	uncovered := RemoveCoveredIntervals(covered)
	fmt.Printf("After removing covered: %v\n\n", uncovered)
}
