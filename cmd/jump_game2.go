package main

import (
	"fmt"
)

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	jumps, farthest, currEnd := 0, 0, 0

	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i]) // Extend reach

		if i == currEnd { // End of current jump
			jumps++
			currEnd = farthest

			if currEnd >= n-1 { // If we reached the last index
				break
			}
		}
	}

	return jumps
}

// Utility function to get max of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // Output: 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // Output: 2
}
