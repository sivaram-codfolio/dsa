package main

import "fmt"

func twoSum(arr []int, target int) []int {
	seen := make(map[int]int, len(arr))

	for i, val := range arr {
		if idx, ok := seen[target-val]; ok {
			return []int{idx, i}
		}
		seen[val] = i
	}

	return nil
}

func main() {
	arr := []int{1, 2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr, target))
}
