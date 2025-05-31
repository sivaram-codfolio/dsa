package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := checkHeight(root)
	return balanced
}

func checkHeight(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := checkHeight(node.Left)
	rightHeight, rightBalanced := checkHeight(node.Right)

	if !leftBalanced || !rightBalanced {
		return 0, false
	}

	if abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return max(leftHeight, rightHeight) + 1, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println("Is the binary tree balanced?")
	fmt.Println(isBalanced(root)) // Output: true
}
