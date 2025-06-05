# Dijkstra's Algorithm

Dijkstra's algorithm is a graph search algorithm that solves the single-source shortest path problem for a graph with non-negative edge weights, producing a shortest path tree.

## Overview

The algorithm maintains a set of unvisited nodes and continuously selects the unvisited node with the smallest tentative distance, marks it as visited, and updates the tentative distances to all its unvisited neighbors.

### Time Complexity
- With Binary Heap: O((V + E) log V)
- With Fibonacci Heap: O(E + V log V)
where V is the number of vertices and E is the number of edges

### Space Complexity
- O(V) for the priority queue and distance array

## Implementation Components

1. **Priority Queue**
   - Binary heap implementation
   - Maintains vertices ordered by distance
   - Supports efficient updates and extractions

2. **Graph Representation**
   - Adjacency list structure
   - Weighted edges
   - Support for both directed and undirected graphs

3. **Core Algorithm**
   - Distance array maintenance
   - Predecessor tracking
   - Path reconstruction

## Usage Examples

```go
// Create a new graph
g := graph.NewGraph(6)

// Add edges (undirected)
g.AddUndirectedEdge(0, 1, 4)  // Edge from 0 to 1 with weight 4
g.AddUndirectedEdge(0, 2, 2)  // Edge from 0 to 2 with weight 2

// Find shortest paths from source vertex 0
distances, predecessors := g.Dijkstra(0)

// Get specific shortest path from 0 to 5
path, distance := g.GetShortestPath(0, 5)
```

## Understanding the Algorithm

### Step-by-Step Process:
1. Initialize distances (infinity for all vertices except source)
2. Create priority queue with source vertex
3. While queue is not empty:
   - Extract vertex with minimum distance
   - Update distances to all neighbors
   - Add updated neighbors to queue

### Example Graph:
```
    4
0 ----→ 1
|     ↗
2    /
|  ↗
↓ /
3
```

### Distance Updates:
```
Initial: [0, ∞, 2, ∞]
After 0: [0, 4, 2, 3]
After 2: [0, 4, 2, 3]
After 3: [0, 4, 2, 3]
After 1: [0, 4, 2, 3]
```

## Applications

1. **Network Routing**
   - IP routing protocols
   - Network link-state algorithms
   - Traffic routing systems

2. **Social Networks**
   - Finding shortest connection paths
   - Influence analysis
   - Network distance calculations

3. **Maps and Navigation**
   - GPS routing
   - Traffic-aware navigation
   - Public transport routing

4. **Games**
   - Pathfinding for NPCs
   - Maze solving
   - Strategy game AI

## Edge Cases Handled

- Disconnected graphs
- No path exists
- Single vertex
- All edges same weight
- Multiple shortest paths

## Performance Optimizations

1. **Priority Queue**
   - Binary heap implementation
   - Efficient decrease-key operation
   - Lazy deletion

2. **Graph Structure**
   - Adjacency list for sparse graphs
   - Efficient edge iteration
   - Memory-efficient storage

3. **Early Termination**
   - Stop when target found
   - Handle unreachable vertices
   - Skip visited nodes

## Common Use Cases

1. **Transportation**
   - Road network navigation
   - Flight routing
   - Logistics optimization

2. **Computer Networks**
   - Packet routing
   - Network flow optimization
   - Load balancing

3. **Resource Planning**
   - Project scheduling
   - Resource allocation
   - Cost minimization

4. **Artificial Intelligence**
   - Pathfinding in games
   - Robot navigation
   - Decision making 