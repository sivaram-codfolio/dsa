package main

import "fmt"

func reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Reverse array ::  ", reverse(arr))
}
