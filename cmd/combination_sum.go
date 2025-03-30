package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(start, sum int, path []int)

	backtrack = func(start, sum int, path []int) {
		if sum == target {
			// Make a copy of path before adding to result
			combination := append([]int{}, path...)
			res = append(res, combination)
			return
		}

		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				continue // Skip if sum exceeds target
			}

			// Choose the number
			path = append(path, candidates[i])
			backtrack(i, sum+candidates[i], path) // Reuse same index
			path = path[:len(path)-1]             // Undo the choice (Backtrack)
		}
	}

	backtrack(0, 0, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // Output: [[2,2,3] [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // Output: [[2,2,2,2] [2,3,3] [3,5]]
}
