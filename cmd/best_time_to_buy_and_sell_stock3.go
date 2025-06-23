package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)

	leftProfits := make([]int, n)
	rightProfits := make([]int, n+1)

	// Left to right: max profit with 1 transaction up to day i
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		leftProfits[i] = max(leftProfits[i-1], prices[i]-minPrice)
	}

	// Right to left: max profit with 1 transaction from day i
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		rightProfits[i] = max(rightProfits[i+1], maxPrice-prices[i])
	}

	// Combine both
	maxProfit := 0
	for i := 0; i < n; i++ {
		maxProfit = max(maxProfit, leftProfits[i]+rightProfits[i])
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	result := maxProfit(prices)
	fmt.Println("Max Profit with at most 2 transactions:", result)
}
