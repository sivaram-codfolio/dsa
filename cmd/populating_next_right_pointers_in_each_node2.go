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

	var curr = root
	var nextHead *Node
	var nextTail *Node

	for curr != nil {
		for curr != nil {
			if curr.Left != nil {
				if nextHead == nil {
					nextHead = curr.Left
					nextTail = curr.Left
				} else {
					nextTail.Next = curr.Left
					nextTail = nextTail.Next
				}
			}

			if curr.Right != nil {
				if nextHead == nil {
					nextHead = curr.Right
					nextTail = curr.Right
				} else {
					nextTail.Next = curr.Right
					nextTail = nextTail.Next
				}
			}

			curr = curr.Next
		}

		curr = nextHead
		nextHead = nil
		nextTail = nil
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
		// Move to the next level's first node
		if level.Left != nil {
			level = level.Left
		} else if level.Right != nil {
			level = level.Right
		} else {
			// Find next level's first node via next pointers
			temp := level.Next
			for temp != nil {
				if temp.Left != nil {
					level = temp.Left
					break
				} else if temp.Right != nil {
					level = temp.Right
					break
				}
				temp = temp.Next
			}
			if temp == nil {
				level = nil
			}
		}
	}
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Right: &Node{Val: 7},
		},
	}

	connect(root)
	fmt.Println("Next pointers per level:")
	printTreeNextPointers(root)
}
