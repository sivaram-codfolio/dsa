package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}

	fmt.Println(maxDepth(root)) // Output: 3
}
