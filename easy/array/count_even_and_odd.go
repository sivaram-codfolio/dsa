package main

import "fmt"

func evenAndOdd(arr []int) (int, int) {
	odd, even := 0, 0
	for _, val := range arr {
		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return odd, even
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	odd, even := evenAndOdd(arr)
	fmt.Println("Odd :: ", odd)
	fmt.Println("even :: ", even)
}
