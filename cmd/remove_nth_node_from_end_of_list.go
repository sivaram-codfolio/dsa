package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy, dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

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

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " → ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Original List: ")
	printList(head)

	n := 2
	head = removeNthFromEnd(head, n)

	fmt.Print("After Removal: ")
	printList(head)
}
