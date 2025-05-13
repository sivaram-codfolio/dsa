package main

import "fmt"

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// If duplicates, skip them
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // Left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0

	fmt.Println("Array:", nums)
	fmt.Println("Target:", target)

	found := search(nums, target)

	if found {
		fmt.Println("Target found in array.")
	} else {
		fmt.Println("Target NOT found in array.")
	}
}
