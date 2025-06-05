package main

import (
	"Breadth_First_Search_Algorithm/bfs"
	"fmt"
)

func main() {
	fmt.Println("=== Breadth-First Search Demonstrations ===\n")

	// Create a sample graph
	/*
	        1
	      /   \
	     2     3
	    / \   / \
	   4   5 6   7
	*/
	g := bfs.NewGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	fmt.Println("Graph Structure:")
	fmt.Println("     1")
	fmt.Println("   /   \\")
	fmt.Println("  2     3")
	fmt.Println(" / \\   / \\")
	fmt.Println("4   5 6   7\n")

	// 1. Simple BFS Traversal
	fmt.Println("1. Simple BFS Traversal from node 1:")
	result := g.SimpleBFS(1)
	fmt.Printf("Visited nodes in order: %v\n\n", result)

	// 2. BFS with Distance
	fmt.Println("2. BFS with Distance from node 1:")
	distances := g.BFSWithDistance(1)
	fmt.Println("Distances from node 1:")
	for node, dist := range distances {
		fmt.Printf("Node %d: %d steps\n", node, dist)
	}
	fmt.Println()

	// 3. Shortest Path
	fmt.Println("3. Shortest Path Examples:")
	paths := [][2]int{{1, 4}, {1, 7}, {4, 7}}
	for _, p := range paths {
		path := g.BFSShortestPath(p[0], p[1])
		fmt.Printf("Shortest path from %d to %d: %v\n", p[0], p[1], path)
	}
	fmt.Println()

	// 4. Level Order Traversal
	fmt.Println("4. Level Order Traversal from node 1:")
	levels := g.BFSLevelOrder(1)
	for i, level := range levels {
		fmt.Printf("Level %d: %v\n", i, level)
	}
	fmt.Println()

	// 5. Multi-Source BFS
	fmt.Println("5. Multi-Source BFS from nodes 1 and 7:")
	sources := []int{1, 7}
	multiDist := g.BFSMultiSource(sources)
	fmt.Println("Distances from nearest source:")
	for node, dist := range multiDist {
		fmt.Printf("Node %d: %d steps\n", node, dist)
	}
	fmt.Println()

	// 6. Demonstrate BFS on a graph with cycles
	fmt.Println("6. BFS on a Cyclic Graph:")
	cyclic := bfs.NewGraph(4)
	cyclic.AddEdge(1, 2)
	cyclic.AddEdge(2, 3)
	cyclic.AddEdge(3, 4)
	cyclic.AddEdge(4, 1) // Creates a cycle

	fmt.Println("\nCyclic Graph Structure:")
	fmt.Println("1 --- 2")
	fmt.Println("|     |")
	fmt.Println("4 --- 3")

	cyclicResult := cyclic.SimpleBFS(1)
	fmt.Printf("BFS traversal order: %v\n", cyclicResult)

	// 7. Demonstrate BFS on a disconnected graph
	fmt.Println("\n7. BFS on a Disconnected Graph:")
	disconnected := bfs.NewGraph(6)
	disconnected.AddEdge(1, 2)
	disconnected.AddEdge(2, 3)
	disconnected.AddEdge(4, 5)
	disconnected.AddEdge(5, 6)

	fmt.Println("\nDisconnected Graph Structure:")
	fmt.Println("1 --- 2 --- 3   4 --- 5 --- 6")

	disconnectedResult := disconnected.SimpleBFS(1)
	fmt.Printf("BFS traversal from node 1: %v\n", disconnectedResult)
	disconnectedResult = disconnected.SimpleBFS(4)
	fmt.Printf("BFS traversal from node 4: %v\n", disconnectedResult)
}
