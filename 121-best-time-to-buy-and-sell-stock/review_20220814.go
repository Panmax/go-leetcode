package main

import "math"

func maxProfit_20220814(prices []int) int {
	minPrice := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}
