package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var build func(left, right int) *TreeNode
	build = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		mid := (left + right) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = build(left, mid-1)
		node.Right = build(mid+1, right)
		return node
	}

	return build(0, len(nums)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println("Inorder Traversal of BST:")
	printInorder(root)
}
