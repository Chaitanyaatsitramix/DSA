package binarytree

import (
	"fmt"
	"strings"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// NewNode creates a new tree node with the given value
func NewNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

// MorrisInorderTraversal performs inorder traversal using Morris algorithm
// Returns the values in inorder sequence
func MorrisInorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	current := root

	for current != nil {
		if current.Left == nil {
			// If no left child, visit current node and move right
			result = append(result, current.Value)
			current = current.Right
		} else {
			// Find the inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create the temporary link
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove the temporary link
				predecessor.Right = nil
				result = append(result, current.Value)
				current = current.Right
			}
		}
	}

	return result
}

// MorrisPreorderTraversal performs preorder traversal using Morris algorithm
func MorrisPreorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Value)
			current = current.Right
		} else {
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				result = append(result, current.Value)
				predecessor.Right = current
				current = current.Left
			} else {
				predecessor.Right = nil
				current = current.Right
			}
		}
	}

	return result
}

// RegularInorderTraversal performs standard recursive inorder traversal
// This is included for comparison with Morris traversal
func RegularInorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorderHelper(root, &result)
	return result
}

// inorderHelper is a recursive helper function for regular inorder traversal
func inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderHelper(node.Left, result)
	*result = append(*result, node.Value)
	inorderHelper(node.Right, result)
}

// BuildExampleTree creates a sample binary tree for testing
//
//	    1
//	   / \
//	  2   3
//	 / \   \
//	4   5   6
func BuildExampleTree() *TreeNode {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Right = NewNode(6)
	return root
}

// BuildComplexTree creates a more complex binary tree for testing
//
//	       1
//	     /   \
//	    2     3
//	   / \   / \
//	  4   5 6   7
//	 / \
//	8   9
func BuildComplexTree() *TreeNode {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)
	root.Left.Left.Left = NewNode(8)
	root.Left.Left.Right = NewNode(9)
	return root
}

// PrintTree prints the binary tree in a simple format
func PrintTree(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	result := "["
	queue := []*TreeNode{root}
	level := 0
	nodesInCurrentLevel := 1
	nodesInNextLevel := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodesInCurrentLevel--

		if node != nil {
			result += fmt.Sprintf("%d ", node.Value)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
			nodesInNextLevel += 2
		} else {
			result += "null "
		}

		if nodesInCurrentLevel == 0 {
			level++
			nodesInCurrentLevel = nodesInNextLevel
			nodesInNextLevel = 0
		}
	}

	return strings.TrimSpace(result) + "]"
}
