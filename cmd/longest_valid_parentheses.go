package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	stack := []int{-1} // Initialize with -1 as a base index
	maxLen := 0

	for i, ch := range s {
		if ch == '(' {
			stack = append(stack, i) // Push the index of '('
		} else {
			stack = stack[:len(stack)-1] // Pop a '(' index
			if len(stack) == 0 {
				stack = append(stack, i) // Store base index for future
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1]) // Compute valid length
			}
		}
	}
	return maxLen
}

// Helper function to get max of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("(()"))    // Output: 2
	fmt.Println(longestValidParentheses(")()())")) // Output: 4
	fmt.Println(longestValidParentheses(""))       // Output: 0
}
