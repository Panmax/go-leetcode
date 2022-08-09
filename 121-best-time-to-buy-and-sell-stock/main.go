package main

import "math"

func maxProfit(prices []int) int {
	profile := 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profile {
			profile = price - minPrice
		}
	}
	return profile
}
