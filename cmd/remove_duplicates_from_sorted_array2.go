package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	j := 2 // Start from the 3rd position

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	newLen := removeDuplicates(nums)
	fmt.Println(newLen)        // Output: 5
	fmt.Println(nums[:newLen]) // Output: [1 1 2 2 3]
}
