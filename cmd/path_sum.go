package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// If it's a leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	remainingSum := targetSum - root.Val

	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
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
	fmt.Printf("Does path sum %d exist? %v\n", targetSum, hasPathSum(root, targetSum))
}
