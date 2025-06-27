package main

import "fmt"

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS helper function
func dfs(node *TreeNode, current int) int {
	if node == nil {
		return 0
	}

	current = current*10 + node.Val

	// If it's a leaf, return the full number
	if node.Left == nil && node.Right == nil {
		return current
	}

	// Recurse left and right
	return dfs(node.Left, current) + dfs(node.Right, current)
}

// Main function
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func main() {
	// Build the tree: [1,2,3]
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}

	result := sumNumbers(root)
	fmt.Println("Sum of Root-to-Leaf Numbers:", result) // Output: 25
}
