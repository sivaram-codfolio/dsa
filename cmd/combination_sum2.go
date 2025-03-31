package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates) // Sort to handle duplicates

	var backtrack func(start int, sum int, path []int)
	backtrack = func(start int, sum int, path []int) {
		if sum == target {
			combination := append([]int{}, path...) // Copy path
			res = append(res, combination)
			return
		}

		for i := start; i < len(candidates); i++ {
			// Skip duplicate numbers at the same depth
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if sum+candidates[i] > target {
				break // Stop early (since array is sorted)
			}

			// Choose number
			path = append(path, candidates[i])
			backtrack(i+1, sum+candidates[i], path) // Move to next index
			path = path[:len(path)-1]               // Undo the choice (Backtrack)
		}
	}

	backtrack(0, 0, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	// Output: [[1,1,6] [1,2,5] [1,7] [2,6]]

	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	// Output: [[1,2,2] [5]]
}
