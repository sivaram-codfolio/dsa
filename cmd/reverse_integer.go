package main

import (
	"fmt"
)

func reverse(x int) int {
	var helper func(int, int) int
	helper = func(x int, rev int) int {
		if x == 0 {
			return rev
		}
		digit := x % 10
		x /= 10

		if rev > 214748364 || (rev == 214748364 && digit > 7) {
			return 0
		}
		if rev < -214748364 || (rev == -214748364 && digit < -8) {
			return 0
		}

		return helper(x, rev*10+digit)
	}

	return helper(x, 0)
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
