package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	results := make([]int, n+1)
	results[0] = 1

	return results
}

func main() {

	digits1 := []int{1, 2, 9}
	fmt.Println(plusOne(digits1))

	digits2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits2))

	digits3 := []int{9}
	fmt.Println(plusOne(digits3))
}
