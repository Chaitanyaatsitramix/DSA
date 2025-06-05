package bfs

// Queue represents a simple FIFO queue
type Queue []interface{}

// Enqueue adds an element to the queue
func (q *Queue) Enqueue(element interface{}) {
	*q = append(*q, element)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue) Dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}

// IsEmpty checks if queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// Export the Graph type and its methods
type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

// NewGraph creates a new graph with given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an undirected edge between vertices v1 and v2
func (g *Graph) AddEdge(v1, v2 int) {
	g.AdjList[v1] = append(g.AdjList[v1], v2)
	g.AdjList[v2] = append(g.AdjList[v2], v1)
}

// SimpleBFS performs basic BFS traversal starting from given vertex
func (g *Graph) SimpleBFS(start int) []int {
	visited := make(map[int]bool)
	result := make([]int, 0)
	queue := Queue{start}
	visited[start] = true

	for !queue.IsEmpty() {
		vertex := queue.Dequeue().(int)
		result = append(result, vertex)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.Enqueue(neighbor)
			}
		}
	}
	return result
}

// BFSWithDistance finds shortest distances from start vertex to all other vertices
func (g *Graph) BFSWithDistance(start int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)
	queue := Queue{start}

	distances[start] = 0
	visited[start] = true

	for !queue.IsEmpty() {
		vertex := queue.Dequeue().(int)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				distances[neighbor] = distances[vertex] + 1
				queue.Enqueue(neighbor)
			}
		}
	}
	return distances
}

// BFSShortestPath finds shortest path between start and end vertices
func (g *Graph) BFSShortestPath(start, end int) []int {
	if start == end {
		return []int{start}
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := Queue{start}
	visited[start] = true

	for !queue.IsEmpty() {
		vertex := queue.Dequeue().(int)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = vertex
				queue.Enqueue(neighbor)

				if neighbor == end {
					// Reconstruct path
					path := []int{end}
					for current := end; current != start; current = parent[current] {
						path = append([]int{parent[current]}, path...)
					}
					return path
				}
			}
		}
	}
	return nil // No path found
}

// BFSLevelOrder performs level-order traversal and returns nodes by levels
func (g *Graph) BFSLevelOrder(start int) [][]int {
	result := make([][]int, 0)
	visited := make(map[int]bool)
	queue := Queue{start}
	visited[start] = true

	for !queue.IsEmpty() {
		levelSize := len(queue)
		currentLevel := make([]int, 0)

		for i := 0; i < levelSize; i++ {
			vertex := queue.Dequeue().(int)
			currentLevel = append(currentLevel, vertex)

			for _, neighbor := range g.AdjList[vertex] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue.Enqueue(neighbor)
				}
			}
		}
		result = append(result, currentLevel)
	}
	return result
}

// BFSMultiSource performs BFS from multiple source vertices simultaneously
func (g *Graph) BFSMultiSource(sources []int) map[int]int {
	distances := make(map[int]int)
	visited := make(map[int]bool)
	queue := make(Queue, 0)

	// Initialize with all source vertices
	for _, source := range sources {
		queue.Enqueue(source)
		distances[source] = 0
		visited[source] = true
	}

	for !queue.IsEmpty() {
		vertex := queue.Dequeue().(int)

		for _, neighbor := range g.AdjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				distances[neighbor] = distances[vertex] + 1
				queue.Enqueue(neighbor)
			}
		}
	}
	return distances
}

// Example usage and test cases will be in bfs_test.go
