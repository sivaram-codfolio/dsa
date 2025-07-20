package main

import (
	"fmt"
	"math"
)

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	if minVal == maxVal {
		return 0
	}

	bucketSize := int(math.Max(1, float64((maxVal-minVal)/(n-1))))
	bucketCount := (maxVal-minVal)/bucketSize + 1

	bucketMin := make([]int, bucketCount)
	bucketMax := make([]int, bucketCount)
	bucketUsed := make([]bool, bucketCount)

	for i := 0; i < bucketCount; i++ {
		bucketMin[i] = math.MaxInt32
		bucketMax[i] = math.MinInt32
	}

	// Place numbers in buckets
	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		if num < bucketMin[idx] {
			bucketMin[idx] = num
		}
		if num > bucketMax[idx] {
			bucketMax[idx] = num
		}
		bucketUsed[idx] = true
	}

	// Compute maximum gap
	prevMax := minVal
	maxGap := 0
	for i := 0; i < bucketCount; i++ {
		if !bucketUsed[i] {
			continue
		}
		if bucketMin[i]-prevMax > maxGap {
			maxGap = bucketMin[i] - prevMax
		}
		prevMax = bucketMax[i]
	}

	return maxGap
}

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println("Maximum Gap:", maximumGap(nums)) // Output: 3
}
