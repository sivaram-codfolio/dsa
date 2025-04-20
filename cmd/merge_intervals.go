package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Sort intervals based on the start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			// Merge if overlapping
			last[1] = max(last[1], curr[1])
		} else {
			// No overlap, add new interval
			merged = append(merged, curr)
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // Output: [[1 6] [8 10] [15 18]]
}
