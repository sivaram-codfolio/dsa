package main

import (
	"fmt"
)

// Function to swap two elements in an array
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// Function to reverse part of the array
func reverse(nums []int, start int) {
	end := len(nums) - 1
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}

// Main function to find the next permutation
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Step 1: Find the first decreasing element from the right
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 2: If such an element exists, find the number to swap with
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j) // Swap the two numbers
	}

	// Step 3: Reverse the remaining part to get the next smallest order
	reverse(nums, i+1)
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1,3,2]

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1,2,3]

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums) // Output: [1,5,1]
}
