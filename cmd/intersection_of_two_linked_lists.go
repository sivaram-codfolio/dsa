package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)

	for headA != nil {
		visited[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if visited[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
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
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	listA := createLinkedList([]int{4, 1})
	listB := createLinkedList([]int{5, 6, 1})
	intersection := createLinkedList([]int{8, 4, 5})

	tempA := listA
	for tempA.Next != nil {
		tempA = tempA.Next
	}
	tempA.Next = intersection

	tempB := listB
	for tempB.Next != nil {
		tempB = tempB.Next
	}
	tempB.Next = intersection

	fmt.Println("List A:")
	printLinkedList(listA)

	fmt.Println("List B:")
	printLinkedList(listB)

	intersectionNode := getIntersectionNode(listA, listB)
	if intersectionNode != nil {
		fmt.Println("Intersection at node with value:", intersectionNode.Val)
	} else {
		fmt.Println("No intersection found.")
	}
}
