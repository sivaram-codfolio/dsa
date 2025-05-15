package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return build(1, n)
}

func build(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var result []*TreeNode

	for i := start; i <= end; i++ {
		leftTrees := build(start, i-1)
		rightTrees := build(i+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				result = append(result, root)
			}
		}
	}

	return result
}

// Helper: Pre-order traversal as string
func preorderTraversal(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return strconv.Itoa(root.Val) + "," + preorderTraversal(root.Left) + "," + preorderTraversal(root.Right)
}

func main() {
	n := 3
	fmt.Println("Generating all unique BSTs for n =", n)

	trees := generateTrees(n)

	fmt.Printf("Total trees generated: %d\n", len(trees))
	for i, tree := range trees {
		fmt.Printf("Tree %d: %s\n", i+1, preorderTraversal(tree))
	}
}
