package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var first, second, prev *TreeNode

func recoverTree(root *TreeNode) {
	first, second, prev = nil, nil, nil
	inorder(root)

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)

	if prev != nil && root.Val < prev.Val {
		if first == nil {
			first = prev
		}
		second = root
	}
	prev = root

	inorder(root.Right)
}

// Helper: Inorder traversal to print tree values
func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Tree: 1 and 3 are swapped in this BST
	// Input:     3
	//           / \
	//          1   4
	//             /
	//            2
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 2},
		},
	}

	fmt.Println("Inorder before recovery:")
	printInorder(root)
	fmt.Println()

	recoverTree(root)

	fmt.Println("Inorder after recovery:")
	printInorder(root)
	fmt.Println()
}
