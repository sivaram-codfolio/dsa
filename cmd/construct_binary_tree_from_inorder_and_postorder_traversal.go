package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	inMap := make(map[int]int)
	for i, val := range inorder {
		inMap[val] = i
	}

	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if inStart > inEnd || postStart > postEnd {
			return nil
		}

		rootVal := postorder[postEnd]
		root := &TreeNode{Val: rootVal}
		inRootIdx := inMap[rootVal]
		leftSize := inRootIdx - inStart

		root.Left = build(inStart, inRootIdx-1, postStart, postStart+leftSize-1)
		root.Right = build(inRootIdx+1, inEnd, postStart+leftSize, postEnd-1)

		return root
	}

	return build(0, len(inorder)-1, 0, len(postorder)-1)
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
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	tree := buildTree(inorder, postorder)
	fmt.Println("Inorder of constructed tree:")
	printTreeInorder(tree)
}
