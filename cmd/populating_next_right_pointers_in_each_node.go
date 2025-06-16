package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftMost := root

	for leftMost.Left != nil {
		head := leftMost
		for head != nil {
			// Connect left -> right
			head.Left.Next = head.Right

			// Connect right -> next left (if exists)
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}

			head = head.Next
		}
		leftMost = leftMost.Left
	}

	return root
}

func printTreeNextPointers(root *Node) {
	level := root
	for level != nil {
		curr := level
		for curr != nil {
			fmt.Print(curr.Val, "->")
			if curr.Next != nil {
				fmt.Print(curr.Next.Val, " ")
			} else {
				fmt.Print("nil ")
			}
			curr = curr.Next
		}
		fmt.Println()
		level = level.Left
	}
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}

	connect(root)
	fmt.Println("Next pointers per level:")
	printTreeNextPointers(root)
}
