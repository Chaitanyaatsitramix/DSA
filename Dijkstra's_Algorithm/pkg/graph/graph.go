package graph

import (
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents a weighted directed graph
type Graph struct {
	Adj [][]Edge
}

// NewGraph creates a new graph with n vertices
func NewGraph(n int) *Graph {
	return &Graph{
		Adj: make([][]Edge, n),
	}
}

// AddEdge adds a directed edge from u to v with weight w
func (g *Graph) AddEdge(from, to, weight int) {
	g.Adj[from] = append(g.Adj[from], Edge{To: to, Weight: weight})
}

// AddUndirectedEdge adds an undirected edge between u and v with weight w
func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// Dijkstra implements Dijkstra's shortest path algorithm
// Returns distances array and predecessors array
func (g *Graph) Dijkstra(start int) ([]int, []int) {
	n := len(g.Adj)
	dist := make([]int, n)
	pred := make([]int, n)
	visited := make([]bool, n)

	// Initialize distances
	for i := range dist {
		dist[i] = math.MaxInt32
		pred[i] = -1
	}
	dist[start] = 0

	// Create priority queue
	pq := make(PriorityQueue, 0)
	startItem := &Item{
		Node:     start,
		Distance: 0,
	}
	pq.Push(startItem)

	for len(pq) > 0 {
		// Get vertex with minimum distance
		u := pq.Pop().(*Item)
		if visited[u.Node] {
			continue
		}
		visited[u.Node] = true

		// Update distances to neighbors
		for _, edge := range g.Adj[u.Node] {
			v := edge.To
			if !visited[v] {
				newDist := dist[u.Node] + edge.Weight
				if newDist < dist[v] {
					dist[v] = newDist
					pred[v] = u.Node
					item := &Item{
						Node:     v,
						Distance: newDist,
					}
					pq.Push(item)
				}
			}
		}
	}

	return dist, pred
}

// GetPath reconstructs the path from start to end using predecessors array
func GetPath(pred []int, end int) []int {
	path := make([]int, 0)
	for curr := end; curr != -1; curr = pred[curr] {
		path = append([]int{curr}, path...)
	}
	return path
}

// GetShortestPath finds the shortest path from start to end
func (g *Graph) GetShortestPath(start, end int) ([]int, int) {
	dist, pred := g.Dijkstra(start)
	path := GetPath(pred, end)
	if len(path) == 0 || path[0] != start {
		return nil, -1 // No path exists
	}
	return path, dist[end]
}
