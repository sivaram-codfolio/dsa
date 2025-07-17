package main

import "fmt"

func merge(arr1, arr2 []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}

func main() {
	arr1 := []int{1, 3, 5}
	arr2 := []int{2, 4, 6}
	fmt.Println("Merge two array by sorted :: ", merge(arr1, arr2))
}
