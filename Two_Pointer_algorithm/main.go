package main

import "fmt"

func main() {
	fmt.Println("=== TWO POINTER ALGORITHM EXAMPLES ===\n")

	fmt.Println("1. OPPOSITE DIRECTION (Convergent) Pointers:")
	fmt.Println("   - Pointers start at opposite ends and move toward each other\n")

	// Call demo from opposite_direction.go
	DemoOppositeDirection()

	fmt.Println("2. SAME DIRECTION (Sliding Window) Pointers:")
	fmt.Println("   - Both pointers move in the same direction at different speeds\n")

	// Call demo from same_direction.go
	DemoSameDirection()
}
