package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rowSet := make([]map[byte]bool, 9)
	colSet := make([]map[byte]bool, 9)
	boxSet := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rowSet[i] = make(map[byte]bool)
		colSet[i] = make(map[byte]bool)
		boxSet[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			boxIndex := (i/3)*3 + j/3

			if rowSet[i][num] || colSet[j][num] || boxSet[boxIndex][num] {
				return false
			}

			rowSet[i][num] = true
			colSet[j][num] = true
			boxSet[boxIndex][num] = true
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board)) // Output: true
}
