package main

import (
	"fmt"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Morris Traversal for Inorder
func inorderTraversalMorris(root *TreeNode) []int {
	var result []int
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// Visit node
			result = append(result, curr.Val)
			curr = curr.Right
		} else {
			// Find inorder predecessor
			predecessor := curr.Left
			for predecessor.Right != nil && predecessor.Right != curr {
				predecessor = predecessor.Right
			}

			// If no thread, create a link
			if predecessor.Right == nil {
				predecessor.Right = curr
				curr = curr.Left
			} else {
				// Remove the thread and visit node
				predecessor.Right = nil
				result = append(result, curr.Val)
				curr = curr.Right
			}
		}
	}

	return result
}

// Main function
func main() {
	// Given tree structure
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}

	// Perform Morris Inorder Traversal
	result := inorderTraversalMorris(root)

	// Print result
	fmt.Println(result) // Output: [1, 3, 2]
}
