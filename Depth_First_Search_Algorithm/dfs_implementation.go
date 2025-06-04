package main

import (
	"fmt"
	"strings"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack implementation for iterative DFS
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

// === Basic DFS Implementations ===

// 1. Recursive DFS (using implicit stack)
func DFSRecursive(graph map[int][]int, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("Visited: %d\n", start)

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			DFSRecursive(graph, neighbor, visited)
		}
	}
}

// 2. Iterative DFS (using explicit stack)
func DFSIterative(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[current] {
			visited[current] = true
			fmt.Printf("Visited: %d\n", current)

			// Add neighbors to stack (in reverse order for consistent traversal)
			neighbors := graph[current]
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}
}

// 3. DFS with custom stack
func DFSWithCustomStack(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	stack := &Stack{}
	stack.Push(start)

	for !stack.IsEmpty() {
		current := stack.Pop()

		if !visited[current] {
			visited[current] = true
			fmt.Printf("Visited: %d\n", current)

			neighbors := graph[current]
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack.Push(neighbors[i])
				}
			}
		}
	}
}

// === Tree Traversals ===

// Pre-order traversal (Root -> Left -> Right)
func PreOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Val)                // Process current
	result = append(result, PreOrder(root.Left)...)  // Left subtree
	result = append(result, PreOrder(root.Right)...) // Right subtree
	return result
}

// In-order traversal (Left -> Root -> Right)
func InOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, InOrder(root.Left)...)  // Left subtree
	result = append(result, root.Val)               // Process current
	result = append(result, InOrder(root.Right)...) // Right subtree
	return result
}

// Post-order traversal (Left -> Right -> Root)
func PostOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, PostOrder(root.Left)...)  // Left subtree
	result = append(result, PostOrder(root.Right)...) // Right subtree
	result = append(result, root.Val)                 // Process current
	return result
}

// Iterative pre-order traversal
func PreOrderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

		// Push right first, then left (so left is processed first)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// === Path Finding ===

// Find path between two nodes
func FindPath(graph map[int][]int, start, target int) []int {
	visited := make(map[int]bool)
	path := []int{}

	if DFSPath(graph, start, target, visited, &path) {
		return path
	}
	return nil
}

func DFSPath(graph map[int][]int, current, target int, visited map[int]bool, path *[]int) bool {
	visited[current] = true
	*path = append(*path, current)

	if current == target {
		return true
	}

	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			if DFSPath(graph, neighbor, target, visited, path) {
				return true
			}
		}
	}

	// Backtrack
	*path = (*path)[:len(*path)-1]
	return false
}

// Find all paths between two nodes
func FindAllPaths(graph map[int][]int, start, target int) [][]int {
	visited := make(map[int]bool)
	path := []int{}
	allPaths := [][]int{}

	DFSAllPaths(graph, start, target, visited, path, &allPaths)
	return allPaths
}

func DFSAllPaths(graph map[int][]int, current, target int, visited map[int]bool, path []int, allPaths *[][]int) {
	visited[current] = true
	path = append(path, current)

	if current == target {
		// Make a copy of current path
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
	} else {
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				DFSAllPaths(graph, neighbor, target, visited, path, allPaths)
			}
		}
	}

	// Backtrack
	visited[current] = false
}

// === Cycle Detection ===

// Detect cycle in undirected graph
func HasCycleUndirected(graph map[int][]int) bool {
	// Convert to undirected graph
	undirectedGraph := makeUndirected(graph)

	visited := make(map[int]bool)

	for node := range undirectedGraph {
		if !visited[node] {
			if DFSCycleUndirected(undirectedGraph, node, -1, visited) {
				return true
			}
		}
	}
	return false
}

func DFSCycleUndirected(graph map[int][]int, current, parent int, visited map[int]bool) bool {
	visited[current] = true

	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			if DFSCycleUndirected(graph, neighbor, current, visited) {
				return true
			}
		} else if neighbor != parent {
			return true // Back edge found
		}
	}
	return false
}

// Detect cycle in directed graph
func HasCycleDirected(graph map[int][]int) bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for node := range graph {
		if !visited[node] {
			if DFSCycleDirected(graph, node, visited, recStack) {
				return true
			}
		}
	}
	return false
}

func DFSCycleDirected(graph map[int][]int, node int, visited, recStack map[int]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			if DFSCycleDirected(graph, neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true // Back edge in recursion stack
		}
	}

	recStack[node] = false
	return false
}

// === Connected Components ===

// Count connected components in an undirected graph
func CountComponents(graph map[int][]int) int {
	// Convert to undirected graph
	undirectedGraph := makeUndirected(graph)

	visited := make(map[int]bool)
	count := 0

	for node := range undirectedGraph {
		if !visited[node] {
			DFSComponent(undirectedGraph, node, visited)
			count++
		}
	}
	return count
}

func DFSComponent(graph map[int][]int, start int, visited map[int]bool) {
	visited[start] = true

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			DFSComponent(graph, neighbor, visited)
		}
	}
}

// Get all connected components as separate slices
func GetConnectedComponents(graph map[int][]int) [][]int {
	// Convert to undirected graph
	undirectedGraph := makeUndirected(graph)

	visited := make(map[int]bool)
	components := [][]int{}

	for node := range undirectedGraph {
		if !visited[node] {
			component := []int{}
			DFSGetComponent(undirectedGraph, node, visited, &component)
			components = append(components, component)
		}
	}
	return components
}

func DFSGetComponent(graph map[int][]int, node int, visited map[int]bool, component *[]int) {
	visited[node] = true
	*component = append(*component, node)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFSGetComponent(graph, neighbor, visited, component)
		}
	}
}

// === Topological Sort ===

func TopologicalSort(graph map[int][]int) []int {
	visited := make(map[int]bool)
	stack := []int{}

	for node := range graph {
		if !visited[node] {
			DFSTopological(graph, node, visited, &stack)
		}
	}

	// Reverse the stack
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

func DFSTopological(graph map[int][]int, node int, visited map[int]bool, stack *[]int) {
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFSTopological(graph, neighbor, visited, stack)
		}
	}

	*stack = append(*stack, node)
}

// === Utility Functions ===

// Helper function to convert directed graph to undirected
func makeUndirected(graph map[int][]int) map[int][]int {
	undirected := make(map[int][]int)

	// Initialize all nodes
	for node := range graph {
		undirected[node] = []int{}
	}

	// Add edges in both directions
	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			// Add edge from node to neighbor
			if !contains(undirected[node], neighbor) {
				undirected[node] = append(undirected[node], neighbor)
			}
			// Add edge from neighbor to node (undirected)
			if _, exists := undirected[neighbor]; !exists {
				undirected[neighbor] = []int{}
			}
			if !contains(undirected[neighbor], node) {
				undirected[neighbor] = append(undirected[neighbor], node)
			}
		}
	}

	return undirected
}

// Helper function to check if slice contains element
func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// Debug DFS with depth tracking
func DFSDebug(graph map[int][]int, start, maxDepth int) {
	visited := make(map[int]bool)
	DFSDebugHelper(graph, start, visited, maxDepth, 0)
}

func DFSDebugHelper(graph map[int][]int, node int, visited map[int]bool, maxDepth, currentDepth int) {
	indent := strings.Repeat("  ", currentDepth)
	fmt.Printf("%sVisiting node %d at depth %d\n", indent, node, currentDepth)

	visited[node] = true

	if currentDepth >= maxDepth {
		fmt.Printf("%sMax depth reached\n", indent)
		return
	}

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			DFSDebugHelper(graph, neighbor, visited, maxDepth, currentDepth+1)
		}
	}
}

// Check if graph is connected (treating as undirected)
func IsConnected(graph map[int][]int) bool {
	if len(graph) == 0 {
		return true
	}

	// Convert to undirected graph
	undirectedGraph := makeUndirected(graph)

	// Get any node as starting point
	var start int
	for node := range undirectedGraph {
		start = node
		break
	}

	visited := make(map[int]bool)
	DFSComponent(undirectedGraph, start, visited)

	// Check if all nodes were visited
	return len(visited) == len(graph)
}

// Calculate maximum depth/height of tree
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main() {
	fmt.Println("=== Depth-First Search Demonstrations ===")

	// Create sample graph
	graph := map[int][]int{
		1: {2, 3},
		2: {4, 5},
		3: {6},
		4: {},
		5: {},
		6: {},
	}

	fmt.Println("Graph:", graph)
	fmt.Println()

	// 1. Recursive DFS
	fmt.Println("1. Recursive DFS from node 1:")
	visited := make(map[int]bool)
	DFSRecursive(graph, 1, visited)
	fmt.Println()

	// 2. Iterative DFS
	fmt.Println("2. Iterative DFS from node 1:")
	DFSIterative(graph, 1)
	fmt.Println()

	// 3. Custom Stack DFS
	fmt.Println("3. DFS with Custom Stack from node 1:")
	DFSWithCustomStack(graph, 1)
	fmt.Println()

	// 4. Tree Traversals
	fmt.Println("4. Tree Traversals:")
	// Build sample tree:    1
	//                      / \
	//                     2   3
	//                    / \
	//                   4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Printf("Pre-order:  %v\n", PreOrder(root))
	fmt.Printf("In-order:   %v\n", InOrder(root))
	fmt.Printf("Post-order: %v\n", PostOrder(root))
	fmt.Printf("Pre-order (Iterative): %v\n", PreOrderIterative(root))
	fmt.Printf("Max Depth: %d\n", MaxDepth(root))
	fmt.Println()

	// 5. Path Finding
	fmt.Println("5. Path Finding:")
	path := FindPath(graph, 1, 6)
	fmt.Printf("Path from 1 to 6: %v\n", path)

	allPaths := FindAllPaths(graph, 1, 6)
	fmt.Printf("All paths from 1 to 6: %v\n", allPaths)
	fmt.Println()

	// 6. Cycle Detection
	fmt.Println("6. Cycle Detection:")
	// Undirected graph with cycle
	cyclicGraph := map[int][]int{
		1: {2, 3},
		2: {1, 3},
		3: {1, 2},
	}
	fmt.Printf("Has cycle (undirected): %t\n", HasCycleUndirected(cyclicGraph))

	// Directed graph with cycle
	directedCyclic := map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}
	fmt.Printf("Has cycle (directed): %t\n", HasCycleDirected(directedCyclic))
	fmt.Println()

	// 7. Connected Components
	fmt.Println("7. Connected Components:")
	disconnectedGraph := map[int][]int{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {},
	}
	fmt.Printf("Number of components: %d\n", CountComponents(disconnectedGraph))
	fmt.Printf("Components: %v\n", GetConnectedComponents(disconnectedGraph))
	fmt.Printf("Is connected: %t\n", IsConnected(graph))
	fmt.Println()

	// 8. Topological Sort
	fmt.Println("8. Topological Sort:")
	dag := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
		0: {},
		1: {},
	}
	topoSort := TopologicalSort(dag)
	fmt.Printf("Topological order: %v\n", topoSort)
	fmt.Println()

	// 9. Debug DFS
	fmt.Println("9. Debug DFS (with depth limit):")
	DFSDebug(graph, 1, 3)
}
