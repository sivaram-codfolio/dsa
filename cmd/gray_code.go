package main

import "fmt"

func grayCode(n int) []int {
	result := []int{0}

	for i := 0; i < n; i++ {
		mask := 1 << i
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]|mask)
		}
	}

	return result
}

func main() {
	fmt.Println(grayCode(2)) // Output: [0, 1, 3, 2]
	fmt.Println(grayCode(3)) // Output: [0, 1, 3, 2, 6, 7, 5, 4]
}
