package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 4, 6, 8, 1, 3, 9, 7}
	fmt.Println("Before Sort", arr)
	bubbleSort(arr)
	fmt.Println("After sort", arr)
}
