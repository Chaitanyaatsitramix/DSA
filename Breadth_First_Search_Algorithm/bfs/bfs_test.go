package bfs

import (
	"reflect"
	"testing"
)

// Helper function to create a test graph
func createTestGraph() *Graph {
	/*
	   Creates this graph:
	        1
	      /   \
	     2     3
	    / \   / \
	   4   5 6   7
	*/
	g := NewGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	return g
}

func TestSimpleBFS(t *testing.T) {
	g := createTestGraph()

	// Test BFS traversal starting from vertex 1
	result := g.SimpleBFS(1)
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SimpleBFS(1) = %v; want %v", result, expected)
	}
}

func TestBFSWithDistance(t *testing.T) {
	g := createTestGraph()

	// Test distances from vertex 1
	distances := g.BFSWithDistance(1)
	expected := map[int]int{
		1: 0, // Distance to self is 0
		2: 1, // One step from 1
		3: 1, // One step from 1
		4: 2, // Two steps from 1 (through 2)
		5: 2, // Two steps from 1 (through 2)
		6: 2, // Two steps from 1 (through 3)
		7: 2, // Two steps from 1 (through 3)
	}

	if !reflect.DeepEqual(distances, expected) {
		t.Errorf("BFSWithDistance(1) = %v; want %v", distances, expected)
	}
}

func TestBFSShortestPath(t *testing.T) {
	g := createTestGraph()

	// Test cases
	tests := []struct {
		start    int
		end      int
		expected []int
	}{
		{1, 4, []int{1, 2, 4}},       // Path: 1 -> 2 -> 4
		{1, 7, []int{1, 3, 7}},       // Path: 1 -> 3 -> 7
		{4, 7, []int{4, 2, 1, 3, 7}}, // Path: 4 -> 2 -> 1 -> 3 -> 7
		{1, 1, []int{1}},             // Same start and end
	}

	for _, test := range tests {
		path := g.BFSShortestPath(test.start, test.end)
		if !reflect.DeepEqual(path, test.expected) {
			t.Errorf("BFSShortestPath(%d, %d) = %v; want %v",
				test.start, test.end, path, test.expected)
		}
	}
}

func TestBFSLevelOrder(t *testing.T) {
	g := createTestGraph()

	// Test level-order traversal starting from vertex 1
	levels := g.BFSLevelOrder(1)
	expected := [][]int{
		{1},          // Level 0
		{2, 3},       // Level 1
		{4, 5, 6, 7}, // Level 2
	}

	if !reflect.DeepEqual(levels, expected) {
		t.Errorf("BFSLevelOrder(1) = %v; want %v", levels, expected)
	}
}

func TestBFSMultiSource(t *testing.T) {
	g := createTestGraph()

	// Test multi-source BFS starting from vertices 1 and 7
	sources := []int{1, 7}
	distances := g.BFSMultiSource(sources)

	expected := map[int]int{
		1: 0, // Distance from source 1
		2: 1, // Distance from 1
		3: 1, // Distance from 1
		4: 2, // Distance from 1
		5: 2, // Distance from 1
		6: 2, // Distance from 1
		7: 0, // Distance from source 7
	}

	if !reflect.DeepEqual(distances, expected) {
		t.Errorf("BFSMultiSource(%v) = %v; want %v", sources, distances, expected)
	}
}

func TestQueue(t *testing.T) {
	var q Queue

	// Test empty queue
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	if elem := q.Dequeue(); elem != nil {
		t.Errorf("Dequeue from empty queue = %v; want nil", elem)
	}

	// Test enqueue and dequeue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueuing")
	}

	expected := []interface{}{1, 2, 3}
	for i, want := range expected {
		if got := q.Dequeue(); got != want {
			t.Errorf("Dequeue %d = %v; want %v", i, got, want)
		}
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func TestDisconnectedGraph(t *testing.T) {
	// Create a disconnected graph
	g := NewGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)

	// Test BFS traversal from vertex 1
	result := g.SimpleBFS(1)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SimpleBFS(1) on disconnected graph = %v; want %v", result, expected)
	}

	// Test distances from vertex 1
	distances := g.BFSWithDistance(1)
	expectedDist := map[int]int{
		1: 0,
		2: 1,
		3: 2,
	}

	if !reflect.DeepEqual(distances, expectedDist) {
		t.Errorf("BFSWithDistance(1) on disconnected graph = %v; want %v",
			distances, expectedDist)
	}
}

func TestCyclicGraph(t *testing.T) {
	// Create a graph with cycles
	g := NewGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1) // Creates a cycle

	// Test BFS traversal
	result := g.SimpleBFS(1)
	expected := []int{1, 2, 4, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SimpleBFS(1) on cyclic graph = %v; want %v", result, expected)
	}

	// Test shortest path in cycle
	path := g.BFSShortestPath(1, 3)
	expectedPath := []int{1, 2, 3}

	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("BFSShortestPath(1, 3) on cyclic graph = %v; want %v",
			path, expectedPath)
	}
}
