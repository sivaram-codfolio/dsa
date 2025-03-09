package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

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

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " → ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	list := createLinkedList([]int{1, 1, 2, 3, 3})
	fmt.Println("Original List:")
	printLinkedList(list)

	result := deleteDuplicates(list)
	fmt.Println("List after removing duplicates:")
	printLinkedList(result)
}
