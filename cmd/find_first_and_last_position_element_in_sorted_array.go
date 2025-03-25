package main

import (
	"fmt"
)

// Helper function to find the first occurrence
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	first := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			first = mid     // Possible first position found
			right = mid - 1 // Move left to find earlier occurrence
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return first
}

// Helper function to find the last occurrence
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	last := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			last = mid     // Possible last position found
			left = mid + 1 // Move right to find later occurrence
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return last
}

// Main function to find the first and last position
func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func main() {
	// Example usage
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target)) // Output: [3, 4]
}
