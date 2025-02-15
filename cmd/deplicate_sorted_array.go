package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("Output: %d, nums1 = %v\n", k1, nums1[:k1])

	nums2 := []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 6, 7}
	k2 := removeDuplicates(nums2)
	fmt.Printf("Output: %d, nums2 = %v\n", k2, nums2[:k2])
}
