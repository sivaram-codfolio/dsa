package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inIndexMap := make(map[int]int)
	for i, val := range inorder {
		inIndexMap[val] = i
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inRootIndex := inIndexMap[rootVal]
		leftSize := inRootIndex - inStart

		root.Left = build(preStart+1, preStart+leftSize, inStart, inRootIndex-1)
		root.Right = build(preStart+leftSize+1, preEnd, inRootIndex+1, inEnd)

		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func printTreeInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printTreeInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printTreeInorder(root.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	tree := buildTree(preorder, inorder)
	fmt.Println("Inorder of constructed tree:")
	printTreeInorder(tree)
}
