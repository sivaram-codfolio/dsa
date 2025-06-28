package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows := len(board)
	cols := len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T' // Temporarily mark safe 'O'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Start from the border
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}
	}

	// Final conversion
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func printBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}

	fmt.Println("Before:")
	printBoard(board)

	solve(board)

	fmt.Println("\nAfter:")
	printBoard(board)
}
