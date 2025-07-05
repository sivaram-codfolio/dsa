package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Insert new nodes after original nodes
	curr := head
	for curr != nil {
		newNode := &Node{Val: curr.Val}
		newNode.Next = curr.Next
		curr.Next = newNode
		curr = newNode.Next
	}

	// Step 2: Assign random pointers
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// Step 3: Separate the two lists
	orig := head
	copyHead := head.Next
	copyCurr := copyHead
	for orig != nil {
		orig.Next = orig.Next.Next
		if copyCurr.Next != nil {
			copyCurr.Next = copyCurr.Next.Next
		}
		orig = orig.Next
		copyCurr = copyCurr.Next
	}

	return copyHead
}

// Helper function to print a list (optional for debugging)
func printList(head *Node) {
	curr := head
	for curr != nil {
		randomVal := "nil"
		if curr.Random != nil {
			randomVal = fmt.Sprintf("%d", curr.Random.Val)
		}
		fmt.Printf("Node{Val: %d, Random: %s} -> ", curr.Val, randomVal)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example list: 1 -> 2, with random pointers
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node1.Next = node2
	node1.Random = node2
	node2.Random = node2

	fmt.Println("Original list:")
	printList(node1)

	copied := copyRandomList(node1)

	fmt.Println("Copied list:")
	printList(copied)
}
