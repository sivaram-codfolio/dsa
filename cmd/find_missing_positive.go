package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// Step 1: Place each number in its correct position if possible
	for i := 0; i < n; {
		correctIdx := nums[i] - 1
		// Swap only if the number is within range [1, n] and not in correct position
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correctIdx] {
			nums[i], nums[correctIdx] = nums[correctIdx], nums[i] // Swap
		} else {
			i++
		}
	}

	// Step 2: Find the first missing positive
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// If all are in correct positions, return n+1
	return n + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))     // Output: 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1})) // Output: 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11})) // Output: 1
}
