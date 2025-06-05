# Topological Sort Algorithm

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u→v, vertex u comes before v in the ordering.

## Overview

A topological sort is only possible for Directed Acyclic Graphs (DAGs). If the graph has a cycle, no valid topological ordering exists.

### Time Complexity
- O(V + E) where V is the number of vertices and E is the number of edges

### Space Complexity
- O(V) for the visited array and recursion stack

## Implementation Components

1. **Graph Representation**
   - Adjacency list structure
   - Directed edges
   - Vertex tracking

2. **Core Algorithm**
   - DFS-based implementation
   - Cycle detection
   - Stack-based result building

3. **Helper Functions**
   - Cycle checking
   - Recursive DFS utility
   - Result construction

## Usage Examples

```go
// Create a new graph
g := graph.NewGraph(4)

// Add directed edges
g.AddEdge(0, 1)  // 0 → 1
g.AddEdge(1, 2)  // 1 → 2
g.AddEdge(2, 3)  // 2 → 3

// Get topological order
order, hasCycle := g.TopologicalSort()
```

## Understanding the Algorithm

### Step-by-Step Process:
1. Create an empty stack
2. For each unvisited vertex:
   - Mark it as visited
   - Recursively process all adjacent vertices
   - Push vertex to stack when done
3. Return stack contents in reverse order

### Example Graph:
```
0 → 1 → 3
    ↓   ↑
    2 ---+
```

### Execution Steps:
```
Visit 0:
  Visit 1:
    Visit 2:
      Visit 3:
        Push 3
      Push 2
    Push 1
  Push 0
Result: [0, 1, 2, 3]
```

## Applications

1. **Build Systems**
   - Dependency resolution
   - Compilation order
   - Package management

2. **Task Scheduling**
   - Project management
   - Job scheduling
   - Resource allocation

3. **Course Prerequisites**
   - Academic planning
   - Learning paths
   - Curriculum design

4. **Data Processing**
   - Pipeline ordering
   - ETL workflows
   - Data dependencies

## Edge Cases Handled

- Cyclic dependencies
- Empty graph
- Single vertex
- Disconnected components
- Multiple valid orderings

## Implementation Details

1. **Graph Structure**
   - Adjacency list representation
   - Efficient edge traversal
   - Memory-efficient storage

2. **Cycle Detection**
   - Recursion stack tracking
   - Early termination
   - Complete cycle reporting

3. **Result Building**
   - Stack-based approach
   - Order preservation
   - Efficient concatenation

## Common Use Cases

1. **Software Build Systems**
   - Library dependencies
   - Compilation order
   - Link order

2. **Project Management**
   - Task dependencies
   - Critical path analysis
   - Resource scheduling

3. **Data Processing**
   - Pipeline stages
   - Data transformations
   - Workflow management

4. **Educational Planning**
   - Course prerequisites
   - Learning paths
   - Skill development

## Algorithm Variations

1. **Kahn's Algorithm**
   - Uses queue instead of stack
   - Iterative approach
   - In-degree tracking

2. **Modified DFS**
   - Timestamps
   - Departure time ordering
   - Path recording

3. **Parallel Processing**
   - Concurrent execution
   - Level-based processing
   - Independent path handling 