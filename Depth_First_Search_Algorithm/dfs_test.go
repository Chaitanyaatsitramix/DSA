package main

import (
	"reflect"
	"testing"
)

// Test helper to create a simple graph
func createTestGraph() map[int][]int {
	return map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}
}

// Test helper to create a tree
func createTestTree() *TreeNode {
	// Tree:    1
	//         / \
	//        2   3
	//       / \
	//      4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	return root
}

// Test Stack implementation
func TestStack(t *testing.T) {
	stack := &Stack{}

	// Test empty stack
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if stack.Size() != 0 {
		t.Error("New stack size should be 0")
	}

	// Test push and pop
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Error("Stack size should be 3")
	}

	if stack.Peek() != 3 {
		t.Error("Peek should return 3")
	}

	if stack.Pop() != 3 {
		t.Error("Pop should return 3")
	}

	if stack.Size() != 2 {
		t.Error("Stack size should be 2 after pop")
	}

	if stack.Pop() != 2 {
		t.Error("Pop should return 2")
	}

	if stack.Pop() != 1 {
		t.Error("Pop should return 1")
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after all pops")
	}
}

// Test Tree Traversals
func TestPreOrder(t *testing.T) {
	root := createTestTree()
	expected := []int{1, 2, 4, 5, 3}
	result := PreOrder(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrder failed. Expected %v, got %v", expected, result)
	}

	// Test empty tree
	if len(PreOrder(nil)) != 0 {
		t.Error("PreOrder of nil should return empty slice")
	}
}

func TestInOrder(t *testing.T) {
	root := createTestTree()
	expected := []int{4, 2, 5, 1, 3}
	result := InOrder(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InOrder failed. Expected %v, got %v", expected, result)
	}

	// Test empty tree
	if len(InOrder(nil)) != 0 {
		t.Error("InOrder of nil should return empty slice")
	}
}

func TestPostOrder(t *testing.T) {
	root := createTestTree()
	expected := []int{4, 5, 2, 3, 1}
	result := PostOrder(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostOrder failed. Expected %v, got %v", expected, result)
	}

	// Test empty tree
	if len(PostOrder(nil)) != 0 {
		t.Error("PostOrder of nil should return empty slice")
	}
}

func TestPreOrderIterative(t *testing.T) {
	root := createTestTree()
	expected := []int{1, 2, 4, 5, 3}
	result := PreOrderIterative(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreOrderIterative failed. Expected %v, got %v", expected, result)
	}

	// Test empty tree
	if len(PreOrderIterative(nil)) != 0 {
		t.Error("PreOrderIterative of nil should return empty slice")
	}
}

func TestMaxDepth(t *testing.T) {
	root := createTestTree()
	expected := 3
	result := MaxDepth(root)

	if result != expected {
		t.Errorf("MaxDepth failed. Expected %d, got %d", expected, result)
	}

	// Test empty tree
	if MaxDepth(nil) != 0 {
		t.Error("MaxDepth of nil should return 0")
	}

	// Test single node
	singleNode := &TreeNode{Val: 1}
	if MaxDepth(singleNode) != 1 {
		t.Error("MaxDepth of single node should return 1")
	}
}

// Test Path Finding
func TestFindPath(t *testing.T) {
	graph := createTestGraph()

	// Test valid path
	path := FindPath(graph, 1, 6)
	expected := []int{1, 3, 6}

	if !reflect.DeepEqual(path, expected) {
		t.Errorf("FindPath failed. Expected %v, got %v", expected, path)
	}

	// Test path to self
	selfPath := FindPath(graph, 1, 1)
	expectedSelf := []int{1}

	if !reflect.DeepEqual(selfPath, expectedSelf) {
		t.Errorf("FindPath to self failed. Expected %v, got %v", expectedSelf, selfPath)
	}

	// Test no path (isolated nodes)
	isolatedGraph := map[int][]int{
		1: {},
		2: {},
	}
	noPath := FindPath(isolatedGraph, 1, 2)

	if noPath != nil {
		t.Errorf("FindPath should return nil for unreachable nodes, got %v", noPath)
	}
}

func TestFindAllPaths(t *testing.T) {
	// Create a graph with multiple paths
	graph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}

	allPaths := FindAllPaths(graph, 1, 4)
	expected := [][]int{{1, 2, 4}, {1, 3, 4}}

	if len(allPaths) != len(expected) {
		t.Errorf("FindAllPaths failed. Expected %d paths, got %d", len(expected), len(allPaths))
	}

	// Check if all expected paths are found
	for _, expectedPath := range expected {
		found := false
		for _, path := range allPaths {
			if reflect.DeepEqual(path, expectedPath) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected path %v not found in result %v", expectedPath, allPaths)
		}
	}
}

// Test Cycle Detection
func TestHasCycleUndirected(t *testing.T) {
	// Graph with cycle
	cyclicGraph := map[int][]int{
		1: {2, 3},
		2: {1, 3},
		3: {1, 2},
	}

	if !HasCycleUndirected(cyclicGraph) {
		t.Error("HasCycleUndirected should detect cycle")
	}

	// Graph without cycle (tree)
	acyclicGraph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {},
		4: {},
		5: {},
	}

	if HasCycleUndirected(acyclicGraph) {
		t.Error("HasCycleUndirected should not detect cycle in tree")
	}

	// Empty graph
	emptyGraph := map[int][]int{}
	if HasCycleUndirected(emptyGraph) {
		t.Error("HasCycleUndirected should not detect cycle in empty graph")
	}
}

func TestHasCycleDirected(t *testing.T) {
	// Directed graph with cycle
	cyclicGraph := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}

	if !HasCycleDirected(cyclicGraph) {
		t.Error("HasCycleDirected should detect cycle")
	}

	// Directed acyclic graph (DAG)
	dagGraph := map[int][]int{
		1: {2, 3},
		2: {4},
		3: {4},
		4: {},
	}

	if HasCycleDirected(dagGraph) {
		t.Error("HasCycleDirected should not detect cycle in DAG")
	}

	// Self-loop
	selfLoop := map[int][]int{
		1: {1},
	}

	if !HasCycleDirected(selfLoop) {
		t.Error("HasCycleDirected should detect self-loop")
	}
}

// Test Connected Components
func TestCountComponents(t *testing.T) {
	// Single connected component
	connectedGraph := createTestGraph()
	if CountComponents(connectedGraph) != 1 {
		t.Error("Connected graph should have 1 component")
	}

	// Multiple components
	disconnectedGraph := map[int][]int{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {},
	}

	expected := 3
	result := CountComponents(disconnectedGraph)
	if result != expected {
		t.Errorf("CountComponents failed. Expected %d, got %d", expected, result)
	}

	// Empty graph
	emptyGraph := map[int][]int{}
	if CountComponents(emptyGraph) != 0 {
		t.Error("Empty graph should have 0 components")
	}
}

func TestGetConnectedComponents(t *testing.T) {
	disconnectedGraph := map[int][]int{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {},
	}

	components := GetConnectedComponents(disconnectedGraph)
	if len(components) != 3 {
		t.Errorf("Expected 3 components, got %d", len(components))
	}

	// Check if all nodes are included
	allNodes := make(map[int]bool)
	for _, component := range components {
		for _, node := range component {
			allNodes[node] = true
		}
	}

	expectedNodes := []int{1, 2, 3, 4, 5}
	for _, node := range expectedNodes {
		if !allNodes[node] {
			t.Errorf("Node %d not found in components", node)
		}
	}
}

func TestIsConnected(t *testing.T) {
	// Connected graph
	connectedGraph := createTestGraph()
	if !IsConnected(connectedGraph) {
		t.Error("Connected graph should return true")
	}

	// Disconnected graph
	disconnectedGraph := map[int][]int{
		1: {2},
		2: {1},
		3: {},
	}

	if IsConnected(disconnectedGraph) {
		t.Error("Disconnected graph should return false")
	}

	// Empty graph
	emptyGraph := map[int][]int{}
	if !IsConnected(emptyGraph) {
		t.Error("Empty graph should be considered connected")
	}
}

// Test Topological Sort
func TestTopologicalSort(t *testing.T) {
	// DAG for testing
	dag := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
		0: {},
		1: {},
	}

	result := TopologicalSort(dag)

	// Verify the result is a valid topological ordering
	position := make(map[int]int)
	for i, node := range result {
		position[node] = i
	}

	// Check that for every edge u -> v, u comes before v
	for node, neighbors := range dag {
		for _, neighbor := range neighbors {
			if position[node] >= position[neighbor] {
				t.Errorf("Invalid topological order: %d should come before %d", node, neighbor)
			}
		}
	}

	// Check that all nodes are included
	if len(result) != len(dag) {
		t.Errorf("Result should contain all nodes. Expected %d, got %d", len(dag), len(result))
	}
}

// Test edge cases
func TestDFSEdgeCases(t *testing.T) {
	// Single node graph
	singleNode := map[int][]int{
		1: {},
	}

	if CountComponents(singleNode) != 1 {
		t.Error("Single node should form 1 component")
	}

	if !IsConnected(singleNode) {
		t.Error("Single node graph should be connected")
	}

	// Graph with self-loops
	selfLoop := map[int][]int{
		1: {1, 2},
		2: {},
	}

	// Should still work correctly
	components := GetConnectedComponents(selfLoop)
	if len(components) != 1 {
		t.Error("Graph with self-loop should have 1 component")
	}
}

// Benchmark tests
func BenchmarkDFSRecursive(b *testing.B) {
	graph := createTestGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		visited := make(map[int]bool)
		DFSRecursive(graph, 1, visited)
	}
}

func BenchmarkDFSIterative(b *testing.B) {
	graph := createTestGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DFSIterative(graph, 1)
	}
}

func BenchmarkTreeTraversal(b *testing.B) {
	root := createTestTree()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PreOrder(root)
		InOrder(root)
		PostOrder(root)
	}
}

func BenchmarkCycleDetection(b *testing.B) {
	graph := createTestGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HasCycleUndirected(graph)
		HasCycleDirected(graph)
	}
}

func BenchmarkConnectedComponents(b *testing.B) {
	disconnectedGraph := map[int][]int{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {},
		6: {7, 8},
		7: {6},
		8: {6},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CountComponents(disconnectedGraph)
	}
}

func BenchmarkPathFinding(b *testing.B) {
	graph := createTestGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FindPath(graph, 1, 6)
	}
}

// Test large graph performance
func TestLargeGraphPerformance(t *testing.T) {
	// Create a larger graph for performance testing
	largeGraph := make(map[int][]int)

	// Create a chain: 1 -> 2 -> 3 -> ... -> 1000
	for i := 1; i < 1000; i++ {
		largeGraph[i] = []int{i + 1}
	}
	largeGraph[1000] = []int{}

	// Test that DFS completes in reasonable time
	start := 1
	target := 1000

	path := FindPath(largeGraph, start, target)
	if path == nil {
		t.Error("Should find path in large graph")
	}

	if len(path) != 1000 {
		t.Errorf("Path length should be 1000, got %d", len(path))
	}

	// Test connected components
	components := CountComponents(largeGraph)
	if components != 1 {
		t.Errorf("Large graph should have 1 component, got %d", components)
	}
}
