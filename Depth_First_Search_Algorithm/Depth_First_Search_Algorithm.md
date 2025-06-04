# Depth-First Search (DFS) Algorithm 🔍📊

## Table of Contents
1. [Overview](#overview)
2. [Core Concept](#core-concept)
3. [Implementation Methods](#implementation-methods)
4. [Stack-Based Approaches](#stack-based-approaches)
5. [Time & Space Complexity](#time--space-complexity)
6. [DFS Variations & Applications](#dfs-variations--applications)
7. [Common Patterns & Problems](#common-patterns--problems)
8. [Step-by-Step Examples](#step-by-step-examples)
9. [Tips and Tricks](#tips-and-tricks)

---

## Overview

**Depth-First Search (DFS)** is a fundamental graph traversal algorithm that explores as far as possible along each branch before backtracking. It uses a **stack** data structure (either explicit or implicit through recursion) to keep track of the path.

### Key Characteristics:
- 🔄 **Strategy**: Go deep first, then backtrack
- 📚 **Data Structure**: Stack (LIFO - Last In, First Out)
- 🎯 **Traversal**: Explores one path completely before trying alternatives
- 🔙 **Backtracking**: Returns to previous node when path is exhausted

### Applications:
- 🗺️ **Pathfinding** in mazes and graphs
- 🌐 **Web crawling** and link analysis
- 🧩 **Puzzle solving** (Sudoku, N-Queens)
- 🔍 **Topological sorting**
- 🌳 **Tree/Graph traversals**
- 🎮 **Game AI** decision trees

---

## Core Concept

### How DFS Works:

1. **Start** at a source node
2. **Mark** current node as visited
3. **Explore** one unvisited neighbor (go deep)
4. **Repeat** step 3 until no unvisited neighbors
5. **Backtrack** to previous node
6. **Continue** until all reachable nodes are visited

### Visual Example - Tree Traversal:
```
        1
       / \
      2   3
     / \   \
    4   5   6

DFS Traversal Order: 1 → 2 → 4 → 5 → 3 → 6

Stack States:
Start: [1]
Visit 1: [2, 3] (add children)
Visit 2: [4, 5, 3] (add children of 2)
Visit 4: [5, 3] (4 has no children)
Visit 5: [3] (5 has no children) 
Visit 3: [6] (add children of 3)
Visit 6: [] (6 has no children)
```

### Graph Example:
```
Graph:     A --- B
           |     |
           |     |
           C --- D

DFS from A: A → B → D → C (one possible order)
```

---

## Implementation Methods

### Method 1: Recursive DFS (Implicit Stack) 🔄

**Uses the call stack** - Most intuitive and commonly used approach.

```go
func DFSRecursive(graph map[int][]int, start int, visited map[int]bool) {
    // Mark current node as visited
    visited[start] = true
    fmt.Printf("Visited: %d\n", start)
    
    // Recursively visit all unvisited neighbors
    for _, neighbor := range graph[start] {
        if !visited[neighbor] {
            DFSRecursive(graph, neighbor, visited)
        }
    }
}

// Wrapper function
func DFS(graph map[int][]int, start int) {
    visited := make(map[int]bool)
    DFSRecursive(graph, start, visited)
}
```

**Advantages:**
- ✅ Simple and intuitive
- ✅ Less code to write
- ✅ Naturally handles backtracking

**Disadvantages:**
- ❌ Stack overflow for deep graphs
- ❌ Limited control over traversal order
- ❌ Uses implicit stack (harder to debug)

### Method 2: Iterative DFS (Explicit Stack) 📚

**Uses an explicit stack** - More control and no recursion limit.

```go
func DFSIterative(graph map[int][]int, start int) {
    visited := make(map[int]bool)
    stack := []int{start}
    
    for len(stack) > 0 {
        // Pop from stack
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if !visited[current] {
            // Mark as visited
            visited[current] = true
            fmt.Printf("Visited: %d\n", current)
            
            // Add all unvisited neighbors to stack
            for _, neighbor := range graph[current] {
                if !visited[neighbor] {
                    stack = append(stack, neighbor)
                }
            }
        }
    }
}
```

**Advantages:**
- ✅ No recursion limit/stack overflow
- ✅ More control over traversal
- ✅ Can pause/resume traversal
- ✅ Easier to debug and trace

**Disadvantages:**
- ❌ More complex code
- ❌ Manual stack management
- ❌ Requires explicit visited tracking

---

## Stack-Based Approaches

### Understanding the Stack Behavior

#### Recursive DFS Stack:
```
Call Stack (Implicit):
dfs(1) → calls dfs(2) → calls dfs(4)
When dfs(4) returns → back to dfs(2) → calls dfs(5)
When dfs(5) returns → back to dfs(2) → back to dfs(1) → calls dfs(3)
```

#### Iterative DFS Stack:
```
Explicit Stack:
[1] → pop 1, push [2,3] → [2,3]
[2,3] → pop 3, push [6] → [2,6] 
[2,6] → pop 6 → [2]
[2] → pop 2, push [4,5] → [4,5]
[4,5] → pop 5 → [4]
[4] → pop 4 → []
```

### Stack Implementation Details

#### Custom Stack for DFS:
```go
type Stack struct {
    items []int
}

func (s *Stack) Push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
    if len(s.items) == 0 {
        return -1 // or panic
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
```

#### DFS with Custom Stack:
```go
func DFSWithCustomStack(graph map[int][]int, start int) {
    visited := make(map[int]bool)
    stack := &Stack{}
    stack.Push(start)
    
    for !stack.IsEmpty() {
        current := stack.Pop()
        
        if !visited[current] {
            visited[current] = true
            fmt.Printf("Visited: %d\n", current)
            
            // Add neighbors in reverse order for correct traversal
            neighbors := graph[current]
            for i := len(neighbors) - 1; i >= 0; i-- {
                if !visited[neighbors[i]] {
                    stack.Push(neighbors[i])
                }
            }
        }
    }
}
```

---

## Time & Space Complexity

| Implementation | Time Complexity | Space Complexity | Notes |
|----------------|----------------|------------------|-------|
| Recursive DFS | O(V + E) | O(V) | V = vertices, E = edges |
| Iterative DFS | O(V + E) | O(V) | Explicit stack space |
| Tree DFS | O(n) | O(h) | n = nodes, h = height |
| Complete Graph | O(V²) | O(V) | Dense graph case |

### Space Complexity Breakdown:
- **Visited Array**: O(V) space
- **Recursion Stack**: O(V) in worst case (linear graph)
- **Explicit Stack**: O(V) in worst case
- **Graph Storage**: O(V + E) for adjacency list

### Why O(V + E)?
- Visit each **vertex** once: O(V)
- Examine each **edge** once: O(E)
- Total: O(V + E)

---

## DFS Variations & Applications

### 1. Tree Traversals 🌳

#### Pre-order Traversal:
```go
func PreOrder(root *TreeNode) {
    if root == nil {
        return
    }
    
    fmt.Printf("%d ", root.Val)  // Process current
    PreOrder(root.Left)          // Recurse left
    PreOrder(root.Right)         // Recurse right
}
```

#### In-order Traversal:
```go
func InOrder(root *TreeNode) {
    if root == nil {
        return
    }
    
    InOrder(root.Left)           // Recurse left
    fmt.Printf("%d ", root.Val)  // Process current
    InOrder(root.Right)          // Recurse right
}
```

#### Post-order Traversal:
```go
func PostOrder(root *TreeNode) {
    if root == nil {
        return
    }
    
    PostOrder(root.Left)         // Recurse left
    PostOrder(root.Right)        // Recurse right
    fmt.Printf("%d ", root.Val)  // Process current
}
```

### 2. Path Finding 🗺️

#### Find Path Between Two Nodes:
```go
func FindPath(graph map[int][]int, start, target int) []int {
    visited := make(map[int]bool)
    path := []int{}
    
    if DFSPath(graph, start, target, visited, &path) {
        return path
    }
    return nil // No path found
}

func DFSPath(graph map[int][]int, current, target int, visited map[int]bool, path *[]int) bool {
    visited[current] = true
    *path = append(*path, current)
    
    if current == target {
        return true // Found target
    }
    
    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            if DFSPath(graph, neighbor, target, visited, path) {
                return true
            }
        }
    }
    
    // Backtrack: remove current node from path
    *path = (*path)[:len(*path)-1]
    return false
}
```

### 3. Cycle Detection 🔄

#### Detect Cycle in Undirected Graph:
```go
func HasCycle(graph map[int][]int) bool {
    visited := make(map[int]bool)
    
    for node := range graph {
        if !visited[node] {
            if DFSCycle(graph, node, -1, visited) {
                return true
            }
        }
    }
    return false
}

func DFSCycle(graph map[int][]int, current, parent int, visited map[int]bool) bool {
    visited[current] = true
    
    for _, neighbor := range graph[current] {
        if !visited[neighbor] {
            if DFSCycle(graph, neighbor, current, visited) {
                return true
            }
        } else if neighbor != parent {
            return true // Back edge found (cycle)
        }
    }
    return false
}
```

### 4. Connected Components 🔗

#### Count Connected Components:
```go
func CountComponents(graph map[int][]int) int {
    visited := make(map[int]bool)
    count := 0
    
    for node := range graph {
        if !visited[node] {
            DFSComponent(graph, node, visited)
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
```

### 5. Topological Sort 📊

#### Topological Ordering (DFS-based):
```go
func TopologicalSort(graph map[int][]int) []int {
    visited := make(map[int]bool)
    stack := []int{}
    
    for node := range graph {
        if !visited[node] {
            DFSTopological(graph, node, visited, &stack)
        }
    }
    
    // Reverse stack to get topological order
    result := make([]int, len(stack))
    for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = stack[j], stack[i]
    }
    return result
}

func DFSTopological(graph map[int][]int, node int, visited map[int]bool, stack *[]int) {
    visited[node] = true
    
    for _, neighbor := range graph[node] {
        if !visited[neighbor] {
            DFSTopological(graph, neighbor, visited, stack)
        }
    }
    
    *stack = append(*stack, node) // Add to stack after processing all neighbors
}
```

---

## Common Patterns & Problems

### Pattern 1: Tree/Graph Traversal 🌳
**Problems:**
- Binary tree traversals
- N-ary tree traversal
- Graph connectivity

### Pattern 2: Path Finding 🗺️
**Problems:**
- Find path between nodes
- All paths from source to target
- Shortest path in unweighted graph

### Pattern 3: Cycle Detection 🔄
**Problems:**
- Detect cycle in graph
- Course schedule problems
- Dependency resolution

### Pattern 4: Connected Components 🔗
**Problems:**
- Number of islands
- Friend circles
- Network connectivity

### Pattern 5: Backtracking 🔙
**Problems:**
- N-Queens problem
- Sudoku solver
- Generate permutations/combinations

### Pattern 6: Flood Fill 🌊
**Problems:**
- Paint bucket tool
- Island problems
- Region marking

---

## Step-by-Step Examples

### Example 1: Tree DFS Traversal
```
Tree:      1
          / \
         2   3
        / \
       4   5

Pre-order DFS: 1 → 2 → 4 → 5 → 3

Step-by-step:
1. Visit 1, add to result: [1]
2. Go left to 2, add to result: [1, 2] 
3. Go left to 4, add to result: [1, 2, 4]
4. Backtrack to 2, go right to 5: [1, 2, 4, 5]
5. Backtrack to 1, go right to 3: [1, 2, 4, 5, 3]
```

### Example 2: Graph DFS with Stack
```
Graph: A-B-C
       |   |
       D---E

Starting from A:

Stack operations:
1. Push A: [A]
2. Pop A, visit A, push neighbors: [B, D]
3. Pop D, visit D, push neighbors: [B, E]
4. Pop E, visit E, push neighbors: [B, C]
5. Pop C, visit C, push neighbors: [B]
6. Pop B, visit B (all neighbors visited): []

Traversal order: A → D → E → C → B
```

### Example 3: Finding Connected Components
```
Graph: 1-2  4-5
       |    |
       3    6

DFS from 1: visits {1, 2, 3} (Component 1)
DFS from 4: visits {4, 5, 6} (Component 2)

Total components: 2
```

---

## Tips and Tricks

### 🎯 **Recursive vs Iterative Choice**
```go
// Use Recursive when:
// - Tree depth is reasonable (< 1000 levels)
// - Code simplicity is important
// - Natural backtracking needed

// Use Iterative when:
// - Deep graphs (avoid stack overflow)
// - Need control over traversal order
// - Memory usage is critical
```

### 🎯 **Handling Different Graph Representations**

#### Adjacency List:
```go
graph := map[int][]int{
    1: {2, 3},
    2: {4, 5},
    3: {6},
}
```

#### Adjacency Matrix:
```go
func DFSMatrix(matrix [][]int, start int) {
    visited := make([]bool, len(matrix))
    DFSMatrixHelper(matrix, start, visited)
}

func DFSMatrixHelper(matrix [][]int, node int, visited []bool) {
    visited[node] = true
    fmt.Printf("Visited: %d\n", node)
    
    for neighbor := 0; neighbor < len(matrix); neighbor++ {
        if matrix[node][neighbor] == 1 && !visited[neighbor] {
            DFSMatrixHelper(matrix, neighbor, visited)
        }
    }
}
```

### 🎯 **Memory Optimization**
```go
// For large graphs, use bitset for visited tracking
type BitSet struct {
    bits []uint64
}

func (bs *BitSet) Set(i int) {
    bs.bits[i/64] |= 1 << (i % 64)
}

func (bs *BitSet) IsSet(i int) bool {
    return bs.bits[i/64]&(1<<(i%64)) != 0
}
```

### 🎯 **Common Pitfalls**
- **Infinite loops**: Always mark nodes as visited
- **Stack overflow**: Use iterative for deep graphs
- **Order dependency**: Be careful about neighbor processing order
- **Disconnected graphs**: Process all unvisited nodes

### 🎯 **Debug Techniques**
```go
func DFSDebug(graph map[int][]int, start int, depth int) {
    visited := make(map[int]bool)
    DFSDebugHelper(graph, start, visited, depth, 0)
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
```

---

## Practice Problems

### Beginner Level 🟢
1. **Binary Tree Traversals** - Pre/In/Post-order
2. **Maximum Depth of Tree** - Basic DFS
3. **Same Tree** - Tree comparison

### Intermediate Level 🟡
1. **Number of Islands** - 2D grid DFS
2. **Path Sum** - Tree path finding
3. **Course Schedule** - Cycle detection
4. **Clone Graph** - Graph traversal

### Advanced Level 🔴
1. **Word Ladder** - Complex graph traversal
2. **Alien Dictionary** - Topological sort
3. **Critical Connections** - Bridge finding
4. **Palindrome Partitioning** - Backtracking

---

## Summary

**Depth-First Search** is a fundamental algorithm that's essential for many graph and tree problems:

### Key Takeaways:
1. **Core Strategy**: Go deep first, then backtrack
2. **Data Structure**: Stack (explicit or implicit via recursion)
3. **Time Complexity**: O(V + E) for graphs, O(n) for trees
4. **Space Complexity**: O(V) for visited tracking + O(h) for stack

### Implementation Choices:
- **Recursive**: Simple, intuitive, but limited by call stack
- **Iterative**: More control, no stack limit, but more complex

### Master These Patterns:
- **Tree Traversals**: Pre/In/Post-order
- **Path Finding**: Source to target
- **Cycle Detection**: Graph validation
- **Connected Components**: Graph partitioning
- **Backtracking**: Solution space exploration

DFS is not just about traversal - it's the foundation for many advanced algorithms including topological sorting, strongly connected components, and complex backtracking problems! 🚀 