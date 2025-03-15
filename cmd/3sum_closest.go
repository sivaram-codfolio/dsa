package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// Step 1: Sort the array
	sort.Ints(nums)
	closestSum := math.MaxInt32 // Large initial value

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Step 2: Update closest sum if needed
			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
			}

			// Step 3: Move pointers
			if sum < target {
				left++ // Increase sum
			} else if sum > target {
				right-- // Decrease sum
			} else {
				return sum // Exact match found
			}
		}
	}
	return closestSum
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // Output: 2
}
