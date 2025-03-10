package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	tree1 := newNode(1)
	tree1.Left = newNode(2)
	tree1.Right = newNode(3)

	tree2 := newNode(1)
	tree2.Left = newNode(2)
	tree2.Right = newNode(3)

	fmt.Println("Are the two trees identical?", isSameTree(tree1, tree2)) // Output: true
}
