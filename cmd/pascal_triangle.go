package main

import (
	"fmt"
)

// Generate Pascal's Triangle
func generate(numRows int) [][]int {
	result := [][]int{}

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1) // Create a row of size (i+1)
		row[0], row[i] = 1, 1   // First and last elements are always 1

		// Compute inner elements
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result = append(result, row)
	}

	return result
}

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}
