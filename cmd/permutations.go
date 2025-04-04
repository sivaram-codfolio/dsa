package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, 0, &res)
	return res
}

func backtrack(nums []int, start int, res *[][]int) {
	// Base case: If all numbers are placed, store the result
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	// Try swapping each element to create permutations
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start] // Swap
		backtrack(nums, start+1, res)               // Recurse
		nums[start], nums[i] = nums[i], nums[start] // Swap back (restore)
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
