package main

import "fmt"

func main() {
	fmt.Println("🔹 KADANE'S ALGORITHM:")
	fmt.Println("1. CLASSIC KADANE'S:")
	fmt.Println("   - Maximum sum of contiguous subarray (O(n) time)\n")
	DemoClassicKadane()

fmt.Println("2. KADANE'S WITH INDICES:")
fmt.Println("   - Also tracks start and end positions of maximum subarray\n")
DemoKadaneWithIndices()

fmt.Println("3. CIRCULAR KADANE'S:")
fmt.Println("   - Maximum sum considering array as circular\n")
DemoCircularKadane()

fmt.Println("4. KADANE'S VARIATIONS:")
fmt.Println("   - Product subarray, constraints, and other modifications\n")
DemoKadaneVariations()
}