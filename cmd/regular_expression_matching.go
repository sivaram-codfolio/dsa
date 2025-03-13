package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true // Empty string matches empty pattern

	// Fill first row (when s is empty)
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1] // Characters match
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] // Ignore '*' and preceding character
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j] // '*' matches multiple times
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("ab", ".*"))                  // true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
}
