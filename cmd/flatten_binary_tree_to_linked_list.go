package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	// Move to the end of the new right subtree and attach the original right subtree
	current := root
	for current.Right != nil {
		current = current.Right
	}
	current.Right = right
}

func printFlattenedTree(root *TreeNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		if root.Left != nil {
			fmt.Print("(Error: Left should be nil) ")
		}
		root = root.Right
	}
	fmt.Println()
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}

	flatten(root)
	fmt.Println("Flattened tree in preorder:")
	printFlattenedTree(root)
}
