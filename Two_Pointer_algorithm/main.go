package main

import "fmt"

func main() {
	fmt.Println("=== ALGORITHM DEMONSTRATIONS ===\n")

	fmt.Println("🔹 TWO POINTER ALGORITHMS:")
	fmt.Println("1. OPPOSITE DIRECTION (Convergent) Pointers:")
	fmt.Println("   - Pointers start at opposite ends and move toward each other\n")
	DemoOppositeDirection()

	fmt.Println("2. SAME DIRECTION (Fast-Slow) Pointers:")
	fmt.Println("   - Both pointers move in the same direction at different speeds\n")
	DemoSameDirection()
}
