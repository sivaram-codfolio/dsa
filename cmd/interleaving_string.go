package main

import "fmt"

func isInterleave(s1, s2, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j > 0 && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

func main() {
	s1 := "aab"
	s2 := "axy"
	s3 := "aaxaby"

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	if isInterleave(s1, s2, s3) {
		fmt.Println("s3 is an interleaving of s1 and s2.")
	} else {
		fmt.Println("s3 is NOT an interleaving of s1 and s2.")
	}
}
