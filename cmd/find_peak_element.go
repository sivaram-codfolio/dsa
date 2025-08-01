package main

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	result := findPeakElement(nums)
	fmt.Println("Peak element index:", result, "Value:", nums[result])
}
