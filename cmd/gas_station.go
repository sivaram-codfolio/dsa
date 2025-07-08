package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		tank += diff
		total += diff

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	result := canCompleteCircuit(gas, cost)
	fmt.Println("Starting gas station index:", result) // Output: 3
}
