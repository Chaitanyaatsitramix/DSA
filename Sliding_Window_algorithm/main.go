package main

import "fmt"

func main() {
	fmt.Println("Sliding Window Algorithm")

	fmt.Println("🔹 SLIDING WINDOW ALGORITHMS:")
	fmt.Println("1. FIXED SIZE SLIDING WINDOW:")
	fmt.Println("   - Window size remains constant, slides across the array\n")
	DemoFixedWindow()

	fmt.Println("2. VARIABLE SIZE SLIDING WINDOW:")
	fmt.Println("   - Window size can expand and contract based on conditions\n")
	DemoVariableWindow()
}
