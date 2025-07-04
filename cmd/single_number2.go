package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	var ones, twos int

	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}

	return ones
}

func main() {
	nums := []int{2, 2, 3, 2}
	result := singleNumber(nums)
	fmt.Println("Single Number:", result)
}
