package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

func main() {
	fmt.Println(myPow(2.0, 10)) // 1024.0
	fmt.Println(myPow(2.1, 3))  // 9.261
	fmt.Println(myPow(2.0, -2)) // 0.25
}
