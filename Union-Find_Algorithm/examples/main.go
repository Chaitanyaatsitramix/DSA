package main

import (
	"fmt"
	"unionfind/pkg/unionfind"
)

func main() {
	fmt.Println("=== Union-Find Algorithm Demonstrations ===\n")

	// Example 1: Array-based Union-Find with Friend Circles
	fmt.Println("Example 1: Friend Circles (Array-based)")
	fmt.Println("Representing friend relationships in a social network")

	// Create Union-Find with 6 people (0-5)
	uf := unionfind.NewArrayUnionFind(6)

	// Add friendships
	fmt.Println("\nAdding friendships:")
	fmt.Println("0 and 1 become friends")
	uf.Union(0, 1)
	fmt.Println("1 and 2 become friends")
	uf.Union(1, 2)
	fmt.Println("3 and 4 become friends")
	uf.Union(3, 4)

	// Check friendships
	fmt.Println("\nChecking relationships:")
	fmt.Printf("Are 0 and 2 in the same friend circle? %v\n", uf.Connected(0, 2))
	fmt.Printf("Are 0 and 4 in the same friend circle? %v\n", uf.Connected(0, 4))
	fmt.Printf("Size of 0's friend circle: %d\n", uf.GetSize(0))

	// Get all friend circles
	fmt.Println("\nAll friend circles:")
	circles := uf.GetSets()
	for i, circle := range circles {
		fmt.Printf("Circle %d: %v\n", i+1, circle)
	}

	// Example 2: Map-based Union-Find with Network Components
	fmt.Println("\nExample 2: Network Components (Map-based)")
	fmt.Println("Representing connected components in a computer network")

	// Create Map-based Union-Find
	network := unionfind.NewMapUnionFind()

	// Add network connections
	fmt.Println("\nAdding network connections:")
	fmt.Println("Connecting 'serverA' to 'serverB'")
	network.Union("serverA", "serverB")
	fmt.Println("Connecting 'serverB' to 'serverC'")
	network.Union("serverB", "serverC")
	fmt.Println("Connecting 'serverD' to 'serverE'")
	network.Union("serverD", "serverE")

	// Add a new server
	fmt.Println("Adding 'serverF' to network")
	network.MakeSet("serverF")

	// Check connections
	fmt.Println("\nChecking network connectivity:")
	fmt.Printf("Are serverA and serverC connected? %v\n", network.Connected("serverA", "serverC"))
	fmt.Printf("Are serverA and serverD connected? %v\n", network.Connected("serverA", "serverD"))
	fmt.Printf("Size of serverA's network: %d\n", network.GetSize("serverA"))

	// Get all network components
	fmt.Println("\nAll network components:")
	components := network.GetSets()
	for root, servers := range components {
		fmt.Printf("Component with root %s: %v\n", root, servers)
	}

	// Example 3: Dynamic Component Growth
	fmt.Println("\nExample 3: Dynamic Component Growth")
	fmt.Println("Demonstrating how components merge and grow")

	// Create a new network for demonstration
	growth := unionfind.NewMapUnionFind()

	// Step 1: Create initial components
	fmt.Println("\nStep 1: Creating initial components")
	growth.Union("A", "B")
	growth.Union("C", "D")
	components = growth.GetSets()
	for root, nodes := range components {
		fmt.Printf("Component with root %s: %v\n", root, nodes)
	}

	// Step 2: Merge components
	fmt.Println("\nStep 2: Merging components")
	growth.Union("B", "C")
	components = growth.GetSets()
	for root, nodes := range components {
		fmt.Printf("Component with root %s: %v\n", root, nodes)
	}

	// Step 3: Add new node to existing component
	fmt.Println("\nStep 3: Adding new node to existing component")
	growth.Union("A", "E")
	components = growth.GetSets()
	for root, nodes := range components {
		fmt.Printf("Component with root %s: %v\n", root, nodes)
	}
}
