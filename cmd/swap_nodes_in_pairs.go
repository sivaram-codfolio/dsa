package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to swap nodes in pairs (iterative)
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// Swap nodes
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// Move pointers
		prev = first
		head = first.Next
	}
	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " → ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	list := createLinkedList([]int{1, 2, 3, 4})
	fmt.Println("Original List:")
	printLinkedList(list)

	result := swapPairs(list)
	fmt.Println("List after swapping pairs:")
	printLinkedList(result)
}
