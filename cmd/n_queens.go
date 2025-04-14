package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	var results [][]string
	board := make([]int, n)
	cols := make(map[int]bool)
	diags1 := make(map[int]bool) // row - col
	diags2 := make(map[int]bool) // row + col

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var boardStr []string
			for i := 0; i < n; i++ {
				rowStr := strings.Repeat(".", n)
				rowStr = rowStr[:board[i]] + "Q" + rowStr[board[i]+1:]
				boardStr = append(boardStr, rowStr)
			}
			results = append(results, boardStr)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diags1[row-col] || diags2[row+col] {
				continue
			}
			board[row] = col
			cols[col], diags1[row-col], diags2[row+col] = true, true, true
			backtrack(row + 1)
			cols[col], diags1[row-col], diags2[row+col] = false, false, false
		}
	}

	backtrack(0)
	return results
}

func main() {
	solutions := solveNQueens(4)
	for _, sol := range solutions {
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
