package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums) // Sort to bring duplicates together
	visited := make([]bool, len(nums))
	var backtrack func(path []int)

	backtrack = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// Skip already used
			if visited[i] {
				continue
			}
			// Skip duplicates
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			backtrack(append(path, nums[i]))
			visited[i] = false
		}
	}

	backtrack([]int{})
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
