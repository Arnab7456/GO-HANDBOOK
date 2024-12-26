package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	bestTime := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if bestTime < prices[i] {
			max := prices[i] - bestTime
			maxProfit = int(math.Max(float64(maxProfit), float64(max)))

		} else {
			bestTime = prices[i]
		}
	}
	return maxProfit
}

func main() {
	// Test case 1
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices)) // Output: 5
}
