package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Last index for merged array
	last := m + n - 1

	// Pointers for nums1 and nums2
	i, j := m-1, n-1

	// Merge from the end
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}

	// Copy remaining elements from nums2 (if any)
	for j >= 0 {
		nums1[last] = nums2[j]
		j--
		last--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // Output: [1, 2, 2, 3, 5, 6]
}
