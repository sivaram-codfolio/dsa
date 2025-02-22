package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x
	var ans int

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(16))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(0))
}
