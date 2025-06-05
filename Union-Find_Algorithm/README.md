# Union-Find (Disjoint Set) Algorithm

Union-Find is a data structure that keeps track of elements partitioned into disjoint (non-overlapping) sets. It provides near-constant-time operations to add new sets, merge existing sets, and determine whether elements are in the same set.

## Overview

The algorithm supports two main operations:
- Union: Merge two sets
- Find: Determine which set an element belongs to

### Time Complexity
- Find: O(α(n)) amortized
- Union: O(α(n)) amortized
where α(n) is the inverse Ackermann function, which grows extremely slowly

### Space Complexity
- O(n) where n is the number of elements

## Implementation Components

1. **Array-based Implementation**
   - Fixed size
   - Integer elements
   - Efficient for known size
   - Memory efficient

2. **HashMap-based Implementation**
   - Dynamic size
   - String/any type elements
   - Flexible element types
   - Good for unknown size

3. **Optimizations**
   - Path compression
   - Union by rank
   - Size tracking

## Usage Examples

```go
// Array-based (for integers)
uf := NewArrayUnionFind(5)
uf.Union(0, 1)
uf.Union(1, 2)
connected := uf.Connected(0, 2) // true

// HashMap-based (for strings)
network := NewMapUnionFind()
network.Union("serverA", "serverB")
network.Union("serverB", "serverC")
connected := network.Connected("serverA", "serverC") // true
```

## Understanding the Algorithm

### Core Operations:

1. **MakeSet(x)**
   - Creates a new set with single element x
   - x is its own parent
   - Initial rank is 0

2. **Find(x)**
   - Returns the root/representative of x's set
   - Performs path compression
   - Makes future operations faster

3. **Union(x, y)**
   - Merges sets containing x and y
   - Uses union by rank
   - Updates size information

### Optimizations:

1. **Path Compression**
   ```
   Before:          After:
   A               A
   ↑               ↑
   B               C
   ↑               ↑
   C               D
   ↑
   D
   ```

2. **Union by Rank**
   ```
   Rank 2         Rank 1
   A              D
   ↑ ↑            ↑
   B C            E
   ```

## Applications

1. **Network Connectivity**
   - Connected components
   - Network topology
   - Peer-to-peer networks

2. **Social Networks**
   - Friend circles
   - Community detection
   - Group membership

3. **Image Processing**
   - Connected components
   - Region labeling
   - Segmentation

4. **Game Development**
   - Team/Alliance systems
   - Territory control
   - Resource sharing

## Edge Cases Handled

- Self-loops
- Empty sets
- Single element sets
- Large components
- Dynamic growth

## Implementation Details

1. **Array-based Features**
   - Fixed-size arrays
   - Integer indices
   - Rank optimization
   - Size tracking

2. **HashMap-based Features**
   - Dynamic growth
   - String keys
   - Flexible types
   - Automatic set creation

3. **Common Features**
   - Path compression
   - Union by rank
   - Component size tracking
   - Set enumeration

## Common Use Cases

1. **Clustering**
   - Data grouping
   - Similarity detection
   - Pattern recognition

2. **Dynamic Connectivity**
   - Network management
   - Resource allocation
   - System monitoring

3. **Percolation**
   - Physical systems
   - Grid connectivity
   - Flow networks

4. **Kruskal's Algorithm**
   - Minimum spanning tree
   - Network design
   - Circuit layout 