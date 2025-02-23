package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt32
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt32
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt32
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt32
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return (math.Max(float64(maxLeftX), float64(maxLeftY)) +
					math.Min(float64(minRightX), float64(minRightY))) / 2
			} else {
				return math.Max(float64(maxLeftX), float64(maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	panic("Input arrays are not sorted correctly")
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("Median:", findMedianSortedArrays(nums1, nums2))

	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	fmt.Println("Median:", findMedianSortedArrays(nums3, nums4))
}
