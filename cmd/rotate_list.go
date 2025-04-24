package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Count length and make list circular
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	tail.Next = head // Make it circular

	// Step 2: Find new tail: (length - k % length - 1) steps from head
	k = k % length
	stepsToNewTail := length - k
	newTail := head
	for i := 1; i < stepsToNewTail; i++ {
		newTail = newTail.Next
	}

	// Step 3: Break circle and return new head
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}

// Helper to create linked list from slice
func createList(values []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range values {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper to print list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	rotated := rotateRight(head, 2)
	printList(rotated) // Output: 4 5 1 2 3
}
