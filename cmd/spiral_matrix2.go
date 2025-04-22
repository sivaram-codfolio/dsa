package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	// Create an n x n matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	for top <= bottom && left <= right {
		// Left to right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Top to bottom
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Right to left
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// Bottom to top
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	matrix := generateMatrix(3)
	for _, row := range matrix {
		fmt.Println(row)
	}
}
