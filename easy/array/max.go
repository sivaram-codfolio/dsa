package main

import "fmt"

func max(arr []int) int {
	max := arr[0]

	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Maximum in array :: ", max(arr))
}
