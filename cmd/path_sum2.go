package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var currentPath []int

	var dfs func(node *TreeNode, remaining int)
	dfs = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}

		currentPath = append(currentPath, node.Val)

		// If it's a leaf and the sum matches
		if node.Left == nil && node.Right == nil && node.Val == remaining {
			pathCopy := make([]int, len(currentPath))
			copy(pathCopy, currentPath)
			result = append(result, pathCopy)
		}

		dfs(node.Left, remaining-node.Val)
		dfs(node.Right, remaining-node.Val)

		// Backtrack
		currentPath = currentPath[:len(currentPath)-1]
	}

	dfs(root, targetSum)
	return result
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}

	targetSum := 22
	paths := pathSum(root, targetSum)
	fmt.Printf("All paths summing to %d:\n", targetSum)
	fmt.Println(paths)
}
