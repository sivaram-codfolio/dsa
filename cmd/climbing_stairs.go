package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first, second := 1, 2
	for i := 3; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}

	return second
}

func main() {
	fmt.Println(climbStairs(3))  // Output: 3
	fmt.Println(climbStairs(5))  // Output: 8
	fmt.Println(climbStairs(10)) // Output: 89
}
