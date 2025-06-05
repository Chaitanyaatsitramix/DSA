# Breadth-First Search (BFS) Algorithm 🌳

## Overview
Breadth-First Search is a graph traversal algorithm that explores a graph level by level, visiting all nodes at the current depth before moving to nodes at the next depth level. It uses a queue data structure to keep track of nodes to visit.

## Key Characteristics
- Explores neighbors before going deeper
- Uses Queue (FIFO - First In First Out)
- Guarantees shortest path in unweighted graphs
- Time Complexity: O(V + E)
- Space Complexity: O(V)

## How BFS Works

### 1. Basic Steps
1. Start at a given source node
2. Add source node to queue
3. While queue is not empty:
   - Dequeue a node
   - Process the node
   - Add all unvisited neighbors to queue
4. Mark nodes as visited to avoid cycles

### 2. Visual Example
```
Graph:
     1
   /   \
  2     3
 / \   / \
4   5 6   7

BFS Order (starting from 1):
Level 0: [1]
Level 1: [2, 3]
Level 2: [4, 5, 6, 7]
```

## Common Applications

### 1. Shortest Path in Unweighted Graphs
- Finding shortest path between two nodes
- Minimum number of moves in a puzzle
- Social networking (finding degrees of connection)

### 2. Level Order Tree Traversal
```
Tree:
     1
   /   \
  2     3
 / \     \
4   5     6

Level Order: [1, 2, 3, 4, 5, 6]
```

### 3. Connected Components
- Finding all connected nodes in a graph
- Network analysis
- Web crawling

## Queue-Based Implementation Methods

### 1. Simple Queue Implementation
```go
type Queue []int

func (q *Queue) Enqueue(n int) {
    *q = append(*q, n)
}

func (q *Queue) Dequeue() int {
    if len(*q) == 0 {
        return -1
    }
    n := (*q)[0]
    *q = (*q)[1:]
    return n
}
```

### 2. BFS with Distance Tracking
```go
func BFS(graph map[int][]int, start int) map[int]int {
    distances := make(map[int]int)
    queue := Queue{start}
    distances[start] = 0
    
    for len(queue) > 0 {
        current := queue.Dequeue()
        for _, neighbor := range graph[current] {
            if _, visited := distances[neighbor]; !visited {
                distances[neighbor] = distances[current] + 1
                queue.Enqueue(neighbor)
            }
        }
    }
    return distances
}
```

## Advanced BFS Variations

### 1. Multi-Source BFS
- Start BFS from multiple nodes simultaneously
- Useful for:
  - Multiple starting points
  - Finding nearest source
  - Flood fill algorithms

### 2. Bidirectional BFS
- Search from both source and target
- Meets in the middle
- Can be faster than regular BFS

### 3. BFS with Layer Counting
- Keep track of levels explicitly
- Useful for:
  - Level-based problems
  - Finding nodes at exact distance
  - Tree level statistics

## Common Problems Solved Using BFS

### 1. Maze Problems
```
Finding shortest path in maze:
[S] [ ] [#] [ ]
[ ] [#] [ ] [ ]
[ ] [ ] [ ] [#]
[#] [ ] [ ] [E]

S = Start, E = End, # = Wall
```

### 2. Word Ladder
- Transform one word to another
- Change one letter at a time
- Example: COLD → CORD → CARD → WARD → WARM

### 3. Network Distance
- Social connections
- Degrees of separation
- Shortest route in transportation

## Best Practices

### 1. Visited Set Management
- Always maintain visited set
- Clear visited set between runs
- Consider using boolean array for dense graphs

### 2. Queue Optimization
- Pre-allocate queue if size known
- Use efficient queue implementation
- Consider deque for bidirectional BFS

### 3. Memory Management
- Clear data structures after use
- Use appropriate data types
- Consider memory-efficient representations

## Time and Space Analysis

### Time Complexity
- Best Case: O(1) - target is start node
- Average Case: O(V + E) - visit all nodes and edges
- Worst Case: O(V + E) - must explore entire graph

### Space Complexity
- Queue: O(W) where W is maximum width
- Visited Set: O(V)
- Total: O(V) in worst case

## Comparison with DFS

### BFS Advantages
1. Finds shortest path in unweighted graphs
2. Better for level-order processing
3. Better for finding closest nodes

### BFS Disadvantages
1. Uses more memory than DFS
2. Not suitable for deep graphs
3. Cannot backtrack easily

## Implementation Tips

### 1. Queue Choice
- Use slice-based queue for small graphs
- Use container/list for large graphs
- Consider custom queue for specific needs

### 2. Graph Representation
- Adjacency list for sparse graphs
- Adjacency matrix for dense graphs
- Custom structure for special cases

### 3. Optimization
- Early termination when target found
- Direction optimization for bidirectional
- Level-based processing for layer needs

## Real-World Applications

### 1. Social Networks
- Friend suggestions
- Degrees of connection
- Community detection

### 2. GPS Navigation
- Finding nearby locations
- Service area mapping
- Route planning

### 3. Web Crawling
- Website indexing
- Link analysis
- Site mapping

## Common Mistakes to Avoid

1. Forgetting to mark nodes as visited
2. Not handling cycles in graph
3. Using wrong data structures
4. Inefficient queue operations
5. Not considering disconnected components

## Practice Problems

1. Find shortest path in binary matrix
2. Word ladder transformation
3. Knight's shortest path on chessboard
4. Network broadcast time
5. Nearest exit from maze

## Conclusion

BFS is a fundamental algorithm that's especially useful when:
- Finding shortest paths
- Level-order processing needed
- Working with social networks
- Exploring state spaces
- Analyzing connectivity

Remember to choose appropriate data structures and consider the specific requirements of your problem when implementing BFS. 