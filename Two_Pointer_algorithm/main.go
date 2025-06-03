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

	fmt.Println("🔹 KADANE'S ALGORITHM:")
	fmt.Println("3. CLASSIC KADANE'S:")
	fmt.Println("   - Maximum sum of contiguous subarray (O(n) time)\n")
	DemoClassicKadane()

	fmt.Println("4. KADANE'S WITH INDICES:")
	fmt.Println("   - Also tracks start and end positions of maximum subarray\n")
	DemoKadaneWithIndices()

	fmt.Println("5. CIRCULAR KADANE'S:")
	fmt.Println("   - Maximum sum considering array as circular\n")
	DemoCircularKadane()

	fmt.Println("6. KADANE'S VARIATIONS:")
	fmt.Println("   - Product subarray, constraints, and other modifications\n")
	DemoKadaneVariations()
}
