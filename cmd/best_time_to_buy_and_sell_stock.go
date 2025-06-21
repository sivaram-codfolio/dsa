package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println("Max Profit:", profit)
}
