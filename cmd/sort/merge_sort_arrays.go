package main

import "fmt"

func bubbleSorting(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	bubbleSorting(arr1)
	bubbleSorting(arr2)
	fmt.Println("arr1 :: ", arr1)
	fmt.Println("arr2 :: ", arr2)

	i, j := 0, 0
	result := make([]int, 0, len(arr1)+len(arr2))

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
	list1 := []int{1, 5, 9}
	list2 := []int{2, 3, 8, 4}
	fmt.Println(mergeSortedArrays(list1, list2)) // Output: [1 2 3 4 5 8 9]
}
