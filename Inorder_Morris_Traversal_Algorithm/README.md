# Morris Traversal Algorithm

Morris Traversal is a tree traversal algorithm that allows us to traverse a binary tree without using recursion or a stack. It achieves O(1) space complexity by temporarily modifying the tree structure during traversal and restoring it afterward.

## Overview

The algorithm works by creating temporary links (threads) from nodes to their inorder successors, allowing us to traverse back up the tree without using a stack or recursion.

### Time Complexity
- Time: O(n) where n is the number of nodes
- Space: O(1) - no extra space needed except for output array

## Implementation Components

1. **Binary Tree Node**
   - Value storage
   - Left child pointer
   - Right child pointer

2. **Core Operations**
   - Morris Inorder Traversal
   - Morris Preorder Traversal
   - Predecessor finding
   - Thread management

3. **Key Features**
   - No recursion
   - No stack
   - Tree restoration
   - Constant space

## Usage Examples

```go
// Create a binary tree
root := NewNode(1)
root.Left = NewNode(2)
root.Right = NewNode(3)

// Perform Morris traversal
result := MorrisInorderTraversal(root)
```

## Understanding the Algorithm

### Step-by-Step Process:

1. **Initialize**
   ```
   current = root
   ```

2. **While current is not null**
   ```
   if current has no left child:
       visit current
       move to right child
   else:
       find predecessor
       if predecessor's right is null:
           link predecessor to current
           move to left child
       else:
           remove link
           visit current
           move to right child
   ```

### Example Tree:
```
     1
    / \
   2   3
  / \   \
 4   5   6
```

### Traversal Steps:
```
1. Start at 1
2. Create thread: 5 -> 1
3. Create thread: 4 -> 2
4. Visit 4
5. Remove thread, visit 2
6. Visit 5
7. Remove thread, visit 1
8. Visit 3
9. Visit 6
```

## Applications

1. **Memory Constrained Systems**
   - Embedded systems
   - Resource-limited devices
   - Real-time systems

2. **Tree Processing**
   - Tree serialization
   - Tree validation
   - Pattern matching

3. **Compiler Design**
   - Expression evaluation
   - Syntax tree processing
   - Code generation

4. **Database Systems**
   - Index traversal
   - Query optimization
   - Tree-based storage

## Edge Cases Handled

- Empty tree
- Single node
- Unbalanced trees
- Complete binary trees
- Linear trees (skewed)

## Implementation Details

1. **Thread Management**
   - Creation of temporary links
   - Safe link removal
   - Tree structure preservation

2. **Traversal Variations**
   - Inorder traversal
   - Preorder traversal
   - Tree modification tracking

3. **Safety Features**
   - Cycle prevention
   - Structure restoration
   - Error handling

## Common Use Cases

1. **Tree Processing**
   - Binary search trees
   - Expression trees
   - Decision trees

2. **Memory Optimization**
   - Large tree traversal
   - Constant space requirements
   - Stack overflow prevention

3. **Algorithm Implementation**
   - Tree serialization
   - Pattern matching
   - Tree comparison

4. **System Design**
   - Memory-efficient systems
   - Real-time processing
   - Embedded systems

## Algorithm Variations

1. **Reverse Morris Traversal**
   - Right-to-left processing
   - Mirror image traversal
   - Backward iteration

2. **Extended Morris Traversal**
   - Multiple passes
   - Custom node processing
   - Additional data collection

3. **Modified Threading**
   - Alternative linking strategies
   - Custom thread management
   - Specialized restoration 