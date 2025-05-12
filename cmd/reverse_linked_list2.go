package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	curr := prev.Next
	var next *ListNode
	for i := 0; i < right-left; i++ {
		next = curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// Helper: Build list from slice
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper: Print list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	left, right := 2, 4

	fmt.Println("Original list:")
	head := buildList(input)
	printList(head)

	result := reverseBetween(head, left, right)

	fmt.Printf("Reversed from position %d to %d:\n", left, right)
	printList(result)
}
