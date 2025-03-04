package main

import "fmt"

func majorityElement(list []int) int {
	candidate, count := 0, 0

	for _, num := range list {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             // Output: 3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // Output: 2
}
