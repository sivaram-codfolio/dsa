package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{}
	afterHead := &ListNode{}
	before := beforeHead
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}

// Helper to print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// Helper to build list from slice
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, v := range vals {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummy.Next
}

func main() {
	head := buildList([]int{1, 4, 3, 2, 5, 2})
	x := 3
	fmt.Print("Original: ")
	printList(head)
	result := partition(head, x)
	fmt.Print("Partitioned: ")
	printList(result)
}
