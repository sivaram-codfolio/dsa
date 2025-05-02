package main

import "fmt"

func combine(n int, k int) [][]int {
	var result [][]int
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		if len(path) == k {
			// Make a copy since slices are shared
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}

		for i := start; i <= n; i++ {
			backtrack(i+1, append(path, i))
		}
	}

	backtrack(1, []int{})
	return result
}

func main() {
	n, k := 4, 2
	fmt.Println(combine(n, k)) // Output: [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
}
