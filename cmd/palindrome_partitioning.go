package main

import (
	"fmt"
)

// Check if a substring is palindrome
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Backtracking function
func backtrack(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		// Reached the end, add a copy of path
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			// Choose
			path = append(path, s[start:end+1])
			// Explore
			backtrack(s, end+1, path, res)
			// Un-choose (backtrack)
			path = path[:len(path)-1]
		}
	}
}

// Main function to call
func partition(s string) [][]string {
	var res [][]string
	backtrack(s, 0, []string{}, &res)
	return res
}

func main() {
	s := "aab"
	result := partition(s)

	fmt.Println("Palindrome partitions of", s, ":")
	for _, partition := range result {
		fmt.Println(partition)
	}
}
