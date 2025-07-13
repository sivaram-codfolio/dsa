package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// Step 1: Detect cycle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// Step 2: Find entry point
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}

func main() {
	// Create list: 3 -> 2 -> 0 -> -4 -> (back to 2)
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // cycle starts at node2

	start := detectCycle(node1)
	if start != nil {
		fmt.Println("Cycle starts at node with value:", start.Val) // Output: 2
	} else {
		fmt.Println("No cycle")
	}
}
