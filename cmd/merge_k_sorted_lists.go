package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MinHeap implementation
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val } // Min-Heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Merge k Sorted Lists using Min Heap
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// Step 1: Push the head of each linked list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	// Step 2: Merge the lists
	dummy := &ListNode{}
	curr := dummy

	for minHeap.Len() > 0 {
		// Extract the smallest node
		minNode := heap.Pop(minHeap).(*ListNode)
		curr.Next = minNode
		curr = curr.Next

		// Push the next node into the heap
		if minNode.Next != nil {
			heap.Push(minHeap, minNode.Next)
		}
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	curr := head
	for _, val := range arr[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head
}

// Helper function to print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	list1 := createList([]int{1, 4, 5})
	list2 := createList([]int{1, 3, 4})
	list3 := createList([]int{2, 6})

	lists := []*ListNode{list1, list2, list3}
	mergedList := mergeKLists(lists)
	printList(mergedList) // Expected Output: 1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6 -> nil
}
