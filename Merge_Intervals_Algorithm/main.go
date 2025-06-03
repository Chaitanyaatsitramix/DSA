package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("🎯 DSA ALGORITHMS DEMONSTRATION")
	fmt.Println(strings.Repeat("=", 40))

	// Merge Intervals Algorithm Demos
	fmt.Println("\n📅 MERGE INTERVALS ALGORITHM")
	DemoBasicMerge()
	DemoInsertInterval()
	DemoMeetingRooms()
	DemoIntervalOperations()

	fmt.Println("✅ All demonstrations completed!")
}