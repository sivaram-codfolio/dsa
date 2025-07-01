package main

import (
	"fmt"
)

// Node structure
type Node struct {
	Val       int
	Neighbors []*Node
}

// Clone graph using DFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)
	return dfs(node, cloned)
}

func dfs(node *Node, cloned map[*Node]*Node) *Node {
	if clone, found := cloned[node]; found {
		return clone
	}

	copy := &Node{Val: node.Val}
	cloned[node] = copy

	for _, neighbor := range node.Neighbors {
		copy.Neighbors = append(copy.Neighbors, dfs(neighbor, cloned))
	}

	return copy
}

// Helper to print graph
func printGraph(node *Node, visited map[*Node]bool) {
	if node == nil || visited[node] {
		return
	}
	visited[node] = true
	fmt.Printf("Node %d: ", node.Val)
	for _, neighbor := range node.Neighbors {
		fmt.Printf("%d ", neighbor.Val)
	}
	fmt.Println()
	for _, neighbor := range node.Neighbors {
		printGraph(neighbor, visited)
	}
}

func main() {
	// Create a test graph: 1-2, 1-4, 2-3, 3-4
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	fmt.Println("Original Graph:")
	printGraph(node1, make(map[*Node]bool))

	cloned := cloneGraph(node1)

	fmt.Println("\nCloned Graph:")
	printGraph(cloned, make(map[*Node]bool))
}
