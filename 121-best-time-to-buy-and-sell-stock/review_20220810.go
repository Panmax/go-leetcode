package main

import "math"

func maxProfit_20220810(prices []int) int {
	profit := 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}
	return profit
}
