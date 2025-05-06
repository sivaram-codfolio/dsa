package main

import "fmt"

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// Fill first row
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// Fill first column
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// Fill the rest of the grid
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid)) // Output: 7
}
