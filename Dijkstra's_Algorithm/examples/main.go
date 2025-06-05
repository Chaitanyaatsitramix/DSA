package main

import (
	"dijkstra/pkg/graph"
	"fmt"
)

func main() {
	fmt.Println("=== Dijkstra's Algorithm Demonstrations ===\n")

	// Example 1: Simple Graph
	/*
	   0 -- 4 -- 1
	   |    |    |
	   3 -- 2 -- 5
	*/
	g1 := graph.NewGraph(6)
	g1.AddUndirectedEdge(0, 4, 1) // 0 -- 4
	g1.AddUndirectedEdge(4, 1, 2) // 4 -- 1
	g1.AddUndirectedEdge(0, 3, 3) // 0 -- 3
	g1.AddUndirectedEdge(3, 2, 2) // 3 -- 2
	g1.AddUndirectedEdge(2, 4, 4) // 2 -- 4
	g1.AddUndirectedEdge(2, 5, 1) // 2 -- 5
	g1.AddUndirectedEdge(1, 5, 3) // 1 -- 5

	fmt.Println("Example 1: Finding shortest paths from vertex 0")
	dist, _ := g1.Dijkstra(0)
	fmt.Println("Distances from source:")
	for i, d := range dist {
		fmt.Printf("To vertex %d: %d\n", i, d)
	}

	// Find path from 0 to 5
	path, distance := g1.GetShortestPath(0, 5)
	fmt.Printf("\nShortest path from 0 to 5: %v (Distance: %d)\n\n", path, distance)

	// Example 2: Complex Graph
	/*
	   0 ----→ 1 ----→ 2
	   ↓       ↓       ↓
	   3 ----→ 4 ----→ 5
	   ↓       ↓       ↓
	   6 ----→ 7 ----→ 8
	*/
	g2 := graph.NewGraph(9)
	// Vertical edges
	g2.AddEdge(0, 3, 2)
	g2.AddEdge(1, 4, 3)
	g2.AddEdge(2, 5, 2)
	g2.AddEdge(3, 6, 4)
	g2.AddEdge(4, 7, 1)
	g2.AddEdge(5, 8, 3)
	// Horizontal edges
	g2.AddEdge(0, 1, 4)
	g2.AddEdge(1, 2, 2)
	g2.AddEdge(3, 4, 3)
	g2.AddEdge(4, 5, 5)
	g2.AddEdge(6, 7, 2)
	g2.AddEdge(7, 8, 4)

	fmt.Println("Example 2: Finding shortest path from vertex 0 to 8")
	path2, distance2 := g2.GetShortestPath(0, 8)
	fmt.Printf("Path: %v\nDistance: %d\n\n", path2, distance2)

	// Example 3: Graph with No Path
	g3 := graph.NewGraph(3)
	g3.AddEdge(0, 1, 1)
	// No path to vertex 2

	fmt.Println("Example 3: Graph with No Path")
	path3, distance3 := g3.GetShortestPath(0, 2)
	if path3 == nil {
		fmt.Println("No path exists from 0 to 2")
	}
	fmt.Printf("Distance: %d\n", distance3)
}
