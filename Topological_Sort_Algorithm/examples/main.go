package main

import (
	"fmt"
	"topological/pkg/graph"
)

func main() {
	fmt.Println("=== Topological Sort Algorithm Demonstrations ===\n")

	// Example 1: Course Prerequisites
	fmt.Println("Example 1: Course Prerequisites")
	fmt.Println("Graph represents course dependencies where edge u->v means course u must be taken before course v")
	/*
	   0 (Intro to Programming) → 2 (Data Structures)    → 3 (Algorithms)
	   ↓                          ↑
	   1 (Discrete Math)     -----+
	*/
	g1 := graph.NewGraph(4)
	g1.AddEdge(0, 2) // Intro to Programming → Data Structures
	g1.AddEdge(1, 2) // Discrete Math → Data Structures
	g1.AddEdge(2, 3) // Data Structures → Algorithms
	g1.AddEdge(0, 1) // Intro to Programming → Discrete Math

	order, hasCycle := g1.TopologicalSort()
	if hasCycle {
		fmt.Println("Error: Course dependencies contain a cycle!")
	} else {
		fmt.Println("Course order:", order)
		fmt.Println("Recommended course sequence:")
		courseNames := []string{
			"Intro to Programming",
			"Discrete Math",
			"Data Structures",
			"Algorithms",
		}
		for i, course := range order {
			fmt.Printf("%d. %s\n", i+1, courseNames[course])
		}
	}

	// Example 2: Build Dependencies
	fmt.Println("\nExample 2: Build Dependencies")
	fmt.Println("Graph represents software build dependencies where edge u->v means package u must be built before package v")
	/*
	   0 (Core) → 1 (UI)     → 3 (Tests)
	              ↑
	   2 (Utils) -+
	*/
	g2 := graph.NewGraph(4)
	g2.AddEdge(0, 1) // Core → UI
	g2.AddEdge(2, 1) // Utils → UI
	g2.AddEdge(1, 3) // UI → Tests

	order2, hasCycle := g2.TopologicalSort()
	if hasCycle {
		fmt.Println("Error: Build dependencies contain a cycle!")
	} else {
		fmt.Println("Build order:", order2)
		fmt.Println("Build sequence:")
		packageNames := []string{
			"Core",
			"UI",
			"Utils",
			"Tests",
		}
		for i, pkg := range order2 {
			fmt.Printf("%d. %s\n", i+1, packageNames[pkg])
		}
	}

	// Example 3: Cyclic Dependencies (Error Case)
	fmt.Println("\nExample 3: Cyclic Dependencies")
	fmt.Println("Graph contains a cycle to demonstrate error handling")
	/*
	   0 → 1 → 2
	   ↑     ↓
	   +-----+
	*/
	g3 := graph.NewGraph(3)
	g3.AddEdge(0, 1)
	g3.AddEdge(1, 2)
	g3.AddEdge(2, 0) // Creates a cycle

	order3, hasCycle := g3.TopologicalSort()
	if hasCycle {
		fmt.Println("Detected cyclic dependencies!")
		fmt.Println("Cannot establish a valid order.")
	} else {
		fmt.Println("Order:", order3)
	}
}
