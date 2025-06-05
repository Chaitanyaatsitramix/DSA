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

## Repository Structure

```
DSA/
├── README.md                           # This comprehensive guide
├── go.mod                             # Go module configuration
├── Binary_Search_Algorithm/           # Day 2 - Search Algorithms
│   ├── Binary_Search_Algorithm.md     # Theoretical guide
│   ├── binary_search.go              # Search implementations
│   └── binary_search_test.go         # Test suite
├── Breadth_First_Search_Algorithm/    # Day 3 - Graph Traversal
│   ├── Breadth_First_Search_Algorithm.md # Algorithm guide
│   ├── bfs/                          # Core BFS package
│   │   ├── bfs.go                    # BFS implementation
│   │   └── bfs_test.go              # BFS tests
│   ├── cmd/                          # Example programs
│   │   └── bfs_demo/
│   │       └── main.go              # Demo program
│   └── go.mod                        # Module configuration
├── Depth_First_Search_Algorithm/      # Day 3 - Graph Traversal
│   ├── Depth_First_Search_Algorithm.md # Algorithm guide
│   ├── dfs_implementation.go         # DFS implementation
│   └── dfs_test.go                   # DFS tests
├── Two_Pointer_algorithm/             # Day 1 - Pointer Techniques
│   ├── Two_Pointer_algorithm.md      # Algorithm guide
│   ├── go.mod                        # Module configuration
│   ├── main.go                       # Example program
│   ├── opposite_direction.go         # Two-pointer implementation
│   └── same_direction.go             # Two-pointer implementation
├── Sliding_Window_algorithm/          # Day 1 - Window Techniques
│   ├── SLIDING_WINDOW.md             # Algorithm guide
│   ├── go.mod                        # Module configuration
│   ├── main.go                       # Example program
│   ├── Fixed_Size_window.go          # Fixed window implementation
│   └── Variable Size_Window.go       # Variable window implementation
├── Merge_Intervals_Algorithm/         # Day 1 - Interval Processing
│   ├── Merge_Intervals_Algorithm.md  # Algorithm guide
│   ├── go.mod                        # Module configuration
│   ├── main.go                       # Example program
│   └── merge_intervals.go            # Implementation
├── Kadane's_Algorithm/               # Day 1 - Dynamic Programming
│   ├── KADANES_ALGORITHM.md          # Algorithm guide
│   ├── go.mod                        # Module configuration
│   ├── main.go                       # Example program
│   └── kadanes_algorithm.go          # Implementation
└── Quickselect_Algorithm/            # Day 3 - Selection Algorithm
    ├── README.md                     # Algorithm guide
    ├── go.mod                        # Module configuration
    ├── quickselect.go               # Core implementation
    └── examples/                    # Example programs
        └── main.go                  # Demo program
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