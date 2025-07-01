package main

import (
	"fmt"
)

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Step 1: Precompute palindrome substrings
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for end := 0; end < n; end++ {
		for start := 0; start <= end; start++ {
			if s[start] == s[end] && (end-start <= 2 || isPalindrome[start+1][end-1]) {
				isPalindrome[start][end] = true
			}
		}
	}

	// Step 2: Compute minimum cuts
	minCuts := make([]int, n)
	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			minCuts[i] = 0
		} else {
			minCuts[i] = i
			for j := 1; j <= i; j++ {
				if isPalindrome[j][i] && minCuts[j-1]+1 < minCuts[i] {
					minCuts[i] = minCuts[j-1] + 1
				}
			}
		}
	}

	return minCuts[n-1]
}

func main() {
	s := "aab"
	fmt.Println("Minimum cuts needed for palindrome partitioning of", s, ":", minCut(s))
}
