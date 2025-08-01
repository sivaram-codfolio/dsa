package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// If start is blocked, no paths
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1

	// Fill first column
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}

	// Fill first row
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	// Fill rest of grid
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid)) // Output: 2
}
