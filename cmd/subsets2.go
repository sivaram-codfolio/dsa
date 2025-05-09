package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // sort to handle duplicates
	res := [][]int{}
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue // skip duplicates
			}
			backtrack(i+1, append(path, nums[i]))
		}
	}

	backtrack(0, []int{})
	return res
}

func main() {
	input := []int{1, 2, 2}
	fmt.Println("Input:", input)

	result := subsetsWithDup(input)

	fmt.Println("Unique subsets:")
	for _, subset := range result {
		fmt.Println(subset)
	}
}
