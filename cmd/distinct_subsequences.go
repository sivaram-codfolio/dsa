package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)

	// dp[i][j] = number of subsequences in s[i:] that match t[j:]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// If t is empty, there's 1 subsequence (empty subsequence)
	for i := 0; i <= m; i++ {
		dp[i][n] = 1
	}

	// Fill dp table from bottom up
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

func main() {
	s := "rabbbit"
	t := "rabbit"

	result := numDistinct(s, t)
	fmt.Printf("Number of distinct subsequences of \"%s\" forming \"%s\": %d\n", s, t, result)
}
