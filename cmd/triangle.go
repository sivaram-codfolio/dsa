package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	n := len(triangle)
	dp := make([]int, n)

	// Initialize with the last row
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	// Bottom-up DP
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + int(math.Min(float64(dp[j]), float64(dp[j+1])))
		}
	}

	return dp[0]
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	result := minimumTotal(triangle)
	fmt.Println("Minimum Path Sum:", result)
}
