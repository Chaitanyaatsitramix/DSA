package graph

// Graph represents a directed graph using adjacency list
type Graph struct {
	V   int     // Number of vertices
	Adj [][]int // Adjacency Lists
}

// NewGraph creates a new graph with V vertices
func NewGraph(v int) *Graph {
	return &Graph{
		V:   v,
		Adj: make([][]int, v),
	}
}

// AddEdge adds a directed edge from v to w
func (g *Graph) AddEdge(v, w int) {
	g.Adj[v] = append(g.Adj[v], w)
}

// TopologicalSort performs topological sorting of the graph
// Returns the sorted vertices and whether the graph has a cycle
func (g *Graph) TopologicalSort() ([]int, bool) {
	// Create a stack for storing the result
	stack := make([]int, 0)

	// Create visited and recursion stack arrays
	visited := make([]bool, g.V)
	recStack := make([]bool, g.V)

	// Call recursive helper function for all unvisited vertices
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			if g.topologicalSortUtil(i, visited, recStack, &stack) {
				return nil, true // Graph has a cycle
			}
		}
	}

	return stack, false
}

// topologicalSortUtil is a recursive helper function for topological sort
// Returns true if graph has a cycle
func (g *Graph) topologicalSortUtil(v int, visited, recStack []bool, stack *[]int) bool {
	// Mark the current node as visited and part of recursion stack
	visited[v] = true
	recStack[v] = true

	// Visit all adjacent vertices
	for _, w := range g.Adj[v] {
		// If vertex is not visited, then recurse on it
		if !visited[w] {
			if g.topologicalSortUtil(w, visited, recStack, stack) {
				return true
			}
		} else if recStack[w] {
			// If the vertex is in recursion stack, then there is a cycle
			return true
		}
	}

	// Remove vertex from recursion stack and add to result
	recStack[v] = false
	*stack = append([]int{v}, *stack...)

	return false
}

// HasCycle checks if the graph contains a cycle
func (g *Graph) HasCycle() bool {
	visited := make([]bool, g.V)
	recStack := make([]bool, g.V)

	// Call recursive helper function for all unvisited vertices
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			if g.hasCycleUtil(i, visited, recStack) {
				return true
			}
		}
	}
	return false
}

// hasCycleUtil is a recursive helper function for cycle detection
func (g *Graph) hasCycleUtil(v int, visited, recStack []bool) bool {
	visited[v] = true
	recStack[v] = true

	for _, w := range g.Adj[v] {
		if !visited[w] {
			if g.hasCycleUtil(w, visited, recStack) {
				return true
			}
		} else if recStack[w] {
			return true
		}
	}

	recStack[v] = false
	return false
}
