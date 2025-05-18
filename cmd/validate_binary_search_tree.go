package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println("Is the tree a valid BST?")
	if isValidBST(root) {
		fmt.Println("✅ Yes")
	} else {
		fmt.Println("❌ No")
	}
}
