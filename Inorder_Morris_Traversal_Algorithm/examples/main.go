package main

import (
	"fmt"
	"morris/pkg/binarytree"
)

func main() {
	fmt.Println("=== Morris Traversal Algorithm Demonstrations ===\n")

	// Example 1: Simple Binary Tree
	fmt.Println("Example 1: Simple Binary Tree")
	fmt.Println("Tree structure:")
	fmt.Println("      1")
	fmt.Println("     / \\")
	fmt.Println("    2   3")
	fmt.Println("   / \\   \\")
	fmt.Println("  4   5   6")

	tree := binarytree.BuildExampleTree()
	fmt.Println("\nTree representation:", binarytree.PrintTree(tree))

	// Perform traversals
	fmt.Println("\nTraversal Results:")
	morrisInorder := binarytree.MorrisInorderTraversal(tree)
	regularInorder := binarytree.RegularInorderTraversal(tree)
	morrisPreorder := binarytree.MorrisPreorderTraversal(tree)

	fmt.Println("Morris Inorder:", morrisInorder)
	fmt.Println("Regular Inorder:", regularInorder)
	fmt.Println("Morris Preorder:", morrisPreorder)

	// Example 2: Complex Binary Tree
	fmt.Println("\nExample 2: Complex Binary Tree")
	fmt.Println("Tree structure:")
	fmt.Println("         1")
	fmt.Println("       /   \\")
	fmt.Println("      2     3")
	fmt.Println("     / \\   / \\")
	fmt.Println("    4   5 6   7")
	fmt.Println("   / \\")
	fmt.Println("  8   9")

	complexTree := binarytree.BuildComplexTree()
	fmt.Println("\nTree representation:", binarytree.PrintTree(complexTree))

	// Perform traversals
	fmt.Println("\nTraversal Results:")
	complexMorrisInorder := binarytree.MorrisInorderTraversal(complexTree)
	complexRegularInorder := binarytree.RegularInorderTraversal(complexTree)
	complexMorrisPreorder := binarytree.MorrisPreorderTraversal(complexTree)

	fmt.Println("Morris Inorder:", complexMorrisInorder)
	fmt.Println("Regular Inorder:", complexRegularInorder)
	fmt.Println("Morris Preorder:", complexMorrisPreorder)

	// Example 3: Step-by-Step Morris Traversal Explanation
	fmt.Println("\nExample 3: Understanding Morris Traversal")
	fmt.Println("Step-by-step process for inorder traversal:")
	fmt.Println("1. Start at root (1)")
	fmt.Println("2. For each node:")
	fmt.Println("   a. If no left child: visit node and go right")
	fmt.Println("   b. If has left child:")
	fmt.Println("      - Find predecessor (rightmost node in left subtree)")
	fmt.Println("      - If predecessor's right is null:")
	fmt.Println("        * Create temporary link to current")
	fmt.Println("        * Go to left child")
	fmt.Println("      - If predecessor's right points to current:")
	fmt.Println("        * Remove temporary link")
	fmt.Println("        * Visit current node")
	fmt.Println("        * Go to right child")

	fmt.Println("\nAdvantages of Morris Traversal:")
	fmt.Println("1. O(1) space complexity - no recursion or stack")
	fmt.Println("2. Preserves original tree structure")
	fmt.Println("3. Useful for memory-constrained systems")
}
