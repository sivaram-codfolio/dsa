package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	longest := 0
	for num := range numMap {
		// Check if it's the start of a sequence
		if !numMap[num-1] {
			currentNum := num
			streak := 1

			for numMap[currentNum+1] {
				currentNum++
				streak++
			}
			if streak > longest {
				longest = streak
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("Longest Consecutive Sequence Length:", longestConsecutive(nums))
}
