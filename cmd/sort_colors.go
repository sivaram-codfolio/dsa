package main

import "fmt"

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			// Swap nums[low] and nums[mid]
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			// Swap nums[mid] and nums[high]
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums) // Output: [0, 0, 1, 1, 2, 2]
}
