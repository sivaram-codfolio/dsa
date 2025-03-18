package main

import (
	"fmt"
)

// ListNode definition
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to reverse k nodes in a linked list
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	// Count the number of nodes
	count := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}

	// Dummy node to handle edge cases
	dummy := &ListNode{0, head}
	prevGroupEnd := dummy

	for count >= k {
		curr := prevGroupEnd.Next
		next := curr.Next

		// Reverse k nodes
		for i := 1; i < k; i++ {
			curr.Next = next.Next
			next.Next = prevGroupEnd.Next
			prevGroupEnd.Next = next
			next = curr.Next
		}

		// Move to the next group
		prevGroupEnd = curr
		count -= k
	}

	return dummy.Next
}

// Helper function to create a linked list from an array
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{arr[0], nil}
	curr := head
	for _, val := range arr[1:] {
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
	}
	return head
}

// Helper function to print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " → ")
		head = head.Next
	}
	fmt.Println("nil")
}

// Main function to test
func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original List: ")
	printList(head)

	k := 2
	head = reverseKGroup(head, k)

	fmt.Print("After k-Reverse: ")
	printList(head)
}
