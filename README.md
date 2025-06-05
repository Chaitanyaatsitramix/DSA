# DSA - Data Structures and Algorithms Learning Journey 🚀

## Overview
This repository documents a comprehensive learning journey through fundamental Data Structures and Algorithms, implemented in Go. Each algorithm includes theoretical explanations, multiple implementation approaches, comprehensive test suites, and performance benchmarks.

## Learning Progression & User Prompts

### Initial Setup
**Prompt 1:** *"create an empty go file"*
- ✅ Created `main.go` with basic Go structure
- ✅ Set up Go module (`go.mod`)

---

### Day 1: Pointer-Based Algorithms

#### 1. Two-Pointer Algorithm
**Prompt 2:** *"Tell me about Two-Pointer Algorithm"*
- ✅ Comprehensive explanation of Two-Pointer technique
- ✅ Implementation of opposite direction and same direction methods
- ✅ Examples: Pair Sum, Remove Duplicates, Container With Most Water
- ✅ Time complexity analysis: O(n), Space: O(1)

#### 2. Sliding Window Algorithm  
**Prompt 3:** *"Tell me about Sliding Window Algorithm"*
- ✅ Detailed guide on Fixed and Variable Sliding Window techniques
- ✅ Implementation of maximum subarray, longest substring problems
- ✅ Step-by-step examples and pattern recognition
- ✅ Applications in substring and subarray problems

#### 3. Merge Intervals Algorithm
**Prompt 4:** *"Tell me about Merge Intervals Algorithm"*
- ✅ Algorithm for merging overlapping intervals
- ✅ Sorting-based approach with O(n log n) complexity
- ✅ Real-world applications in scheduling and time management

#### 4. Kadane's Algorithm
**Prompt 5:** *"Tell me about Kadane's Algorithm"*
- ✅ Maximum subarray sum problem solution
- ✅ Dynamic programming approach
- ✅ O(n) time complexity implementation
- ✅ Applications in array processing

---

### Day 2: Search Algorithms

#### 5. Binary Search Algorithm
**Prompt 6:** *"tell me about Binary Search Algorithm"*
- ✅ Comprehensive guide with variations:
  - Standard Binary Search (iterative & recursive)
  - Lower Bound & Upper Bound
  - Search Range & Rotated Array Search
  - Peak Element & Square Root algorithms
- ✅ Complete implementations in `binary_search.go`
- ✅ Test suite with edge cases
- ✅ Performance benchmarks

---

### Day 3: Graph and Selection Algorithms

#### 6. Depth-First Search (DFS) Algorithm
**Prompt 7:** *"tell me about Depth-First Search Algorithm"*
- ✅ Comprehensive DFS guide covering:
  - Stack-based implementations
  - Tree traversals
  - Path finding algorithms
  - Cycle detection
  - Connected components
  - Topological sorting

#### 7. Breadth-First Search (BFS) Algorithm
**Prompt 8:** *"tell me about Breadth First Search Algorithm"*
- ✅ Complete BFS implementation with:
  - Queue-based approach
  - Level-order traversal
  - Shortest path finding
  - Distance tracking
  - Multi-source BFS
- ✅ Practical examples and applications

#### 8. Quickselect Algorithm
**Prompt 9:** *"explain me what is Quickselect_Algorithm and its methods in Array with example"*
- ✅ Comprehensive implementation with:
  - Basic Quickselect with O(n) average case complexity
  - Median of Medians for guaranteed O(n) worst case
  - Lomuto's partitioning scheme
  - Edge case handling
- ✅ Example applications:
  - Finding kth smallest element
  - Finding median in unsorted array
  - Handling arrays with duplicates

#### 9. Morris Traversal Algorithm
**Prompt 10:** *"explain me what is Inorder_Morris_Traversal_Algorithm and its methods in Binary Tree with example"*
- ✅ Space-efficient binary tree traversal:
  - O(1) space complexity without recursion/stack
  - Inorder traversal implementation
  - Preorder traversal variation
  - Thread creation and management
- ✅ Example implementations:
  - Simple binary tree traversal
  - Complex tree processing
  - Tree structure preservation
- ✅ Applications:
  - Memory-constrained systems
  - Large tree processing
  - Embedded systems

#### 10. KMP Algorithm
**Prompt 11:** *"explain me what is KMP Algorithm and its methods with example"*
- ✅ Pattern matching algorithm
- ✅ Implementation details
- ✅ Example applications

#### 11. Dijkstra's Algorithm
**Prompt 12:** *"explain me what is Dijkstra's Algorithm and its methods with example"*
- ✅ Shortest path algorithm
- ✅ Priority queue implementation
- ✅ Graph traversal applications

#### 12. Topological Sort Algorithm
**Prompt 13:** *"explain me what is Topological Sort Algorithm and its methods with example"*
- ✅ Directed acyclic graph sorting
- ✅ Dependency resolution
- ✅ Application examples

#### 13. Trie Search Insert Algorithm
**Prompt 14:** *"explain me what is Trie Search Insert Algorithm and its methods with example"*
- ✅ Prefix tree implementation
- ✅ Search and insert operations
- ✅ String processing applications

#### 14. Union-Find Algorithm
**Prompt 15:** *"explain me what is Union-Find Algorithm and its methods with example"*
- ✅ Disjoint set data structure
- ✅ Union and find operations
- ✅ Connected components applications

## Repository Structure

```
DSA/
├── README.md                           # This comprehensive guide
├── go.mod                             # Go module configuration
├── Binary_Search_Algorithm/           # Search Algorithms
│   ├── Binary_Search_Algorithm.md     # Algorithm guide
│   ├── binary_search.go              # Implementation
│   └── binary_search_test.go         # Test suite
├── Breadth_First_Search_Algorithm/    # Graph Traversal
│   ├── Breadth_First_Search_Algorithm.md
│   ├── bfs/
│   │   ├── bfs.go
│   │   └── bfs_test.go
│   ├── cmd/
│   │   └── bfs_demo/
│   │       └── main.go
│   └── go.mod
├── Depth_First_Search_Algorithm/      # Graph Traversal
│   ├── Depth_First_Search_Algorithm.md
│   ├── dfs_implementation.go
│   └── dfs_test.go
├── Dijkstra's_Algorithm/             # Shortest Path
│   ├── README.md
│   ├── go.mod
│   ├── examples/
│   │   └── main.go
│   └── pkg/
│       └── graph/
│           ├── graph.go
│           └── priority_queue.go
├── Inorder_Morris_Traversal_Algorithm/ # Tree Traversal
│   ├── README.md
│   ├── go.mod
│   ├── examples/
│   │   └── main.go
│   └── pkg/
│       └── binarytree/
│           └── binary_tree.go
├── Kadane's_Algorithm/               # Dynamic Programming
│   ├── KADANES_ALGORITHM.md
│   ├── go.mod
│   ├── main.go
│   └── kadanes_algorithm.go
├── KMP_Algorithm/                    # String Matching
│   ├── README.md
│   ├── go.mod
│   ├── kmp.go
│   └── examples/
│       └── main.go
├── Merge_Intervals_Algorithm/        # Interval Processing
│   ├── Merge_Intervals_Algorithm.md
│   ├── go.mod
│   ├── main.go
│   └── merge_intervals.go
├── Quickselect_Algorithm/           # Selection Algorithm
│   ├── README.md
│   ├── go.mod
│   ├── quickselect.go
│   └── examples/
│       └── main.go
├── Sliding_Window_algorithm/        # Window Techniques
│   ├── SLIDING_WINDOW.md
│   ├── go.mod
│   ├── main.go
│   ├── Fixed_Size_window.go
│   └── Variable_Size_Window.go
├── Topological_Sort_Algorithm/      # Graph Algorithm
│   ├── README.md
│   ├── go.mod
│   ├── examples/
│   │   └── main.go
│   └── pkg/
│       └── graph/
│           └── graph.go
├── Trie_Search_Insert_Algorithm/    # Tree Data Structure
│   ├── README.md
│   ├── go.mod
│   ├── examples/
│   │   └── main.go
│   └── pkg/
│       └── trie/
│           └── trie.go
├── Two_Pointer_algorithm/          # Pointer Techniques
│   ├── Two_Pointer_algorithm.md
│   ├── go.mod
│   ├── main.go
│   ├── opposite_direction.go
│   └── same_direction.go
└── Union-Find_Algorithm/          # Disjoint Sets
    ├── README.md
    ├── go.mod
    ├── examples/
    │   └── main.go
    └── pkg/
        └── unionfind/
            ├── array_unionfind.go
            └── map_unionfind.go
```

## Key Learning Outcomes

### 🎯 **Algorithm Mastery**
- **Binary Search**: O(log n) search efficiency with variations
- **DFS & BFS**: Graph traversal with multiple applications
- **Two-Pointer**: O(n) linear scanning techniques
- **Sliding Window**: Efficient substring/subarray problem solving
- **Merge Intervals**: O(n log n) interval processing
- **Kadane's Algorithm**: O(n) dynamic programming approach
- **Quickselect**: O(n) selection algorithm with deterministic variant
- **Morris Traversal**: O(1) space tree traversal without recursion
- **KMP Algorithm**: O(n + m) string pattern matching
- **Dijkstra's Algorithm**: O((V + E)logV) shortest path finding
- **Topological Sort**: O(V + E) DAG ordering
- **Trie**: O(L) prefix tree operations
- **Union-Find**: Near O(1) disjoint set operations

### 🧪 **Testing Excellence**
- Comprehensive test coverage
- Edge case handling
- Performance benchmarking
- Large dataset validation

### ⚡ **Performance Insights**
- Time complexity analysis for each algorithm
- Space complexity considerations
- Optimization techniques
- Trade-off analysis

## Usage Instructions

### Running Tests
```bash
# Test specific algorithm
cd Algorithm_Directory && go test -v

# Run benchmarks
go test -bench=.

# Test with coverage
go test -cover
```

### Running Demonstrations
```bash
# Run any algorithm example
cd Algorithm_Directory && go run .
```

## Learning Philosophy
This repository demonstrates a systematic approach to algorithm learning:
1. Theoretical Understanding
2. Multiple Implementations
3. Comprehensive Testing
4. Performance Analysis
5. Real-world Applications

---

*Happy Coding! 🚀*