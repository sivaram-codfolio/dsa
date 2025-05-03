package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		// Add a copy of path to result
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			backtrack(i+1, append(path, nums[i]))
		}
	}

	backtrack(0, []int{})
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
