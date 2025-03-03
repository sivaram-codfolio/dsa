package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return (1 << 31) - 1
	}

	negative := (dividend < 0) != (divisor < 0)
	dividend, divisor = abs(dividend), abs(divisor)

	quotient := 0
	for dividend >= divisor {
		temp, multiple := divisor, 1

		for dividend >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}

		dividend -= temp
		quotient += multiple
	}

	if negative {
		quotient = -quotient
	}

	if quotient < -(1 << 31) {
		return -(1 << 31)
	} else if quotient > (1<<31)-1 {
		return (1 << 31) - 1
	}

	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(divide(10, 3))           // Output: 3
	fmt.Println(divide(7, -3))           // Output: -2
	fmt.Println(divide(0, 1))            // Output: 0
	fmt.Println(divide(1, 1))            // Output: 1
	fmt.Println(divide(-2147483648, -1)) // Output: 2147483647 (overflow case)
}
