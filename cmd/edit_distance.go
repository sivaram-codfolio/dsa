package main

import "fmt"

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i // delete all characters
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // insert all characters
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],   // delete
					dp[i][j-1],   // insert
					dp[i-1][j-1], // replace
				)
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))           // Output: 3
	fmt.Println(minDistance("intention", "execution")) // Output: 5
}
