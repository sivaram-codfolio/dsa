package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}

	dfs(root)
	return result
}

func main() {
	// Example Tree:   1
	//                   \
	//                    2
	//                   /
	//                  3
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	result := postorderTraversal(root)
	fmt.Println("Postorder Traversal:", result) // Output: [3, 2, 1]
}
