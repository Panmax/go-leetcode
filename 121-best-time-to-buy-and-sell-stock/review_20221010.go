package main

import "math"

func maxProfit_20221010(prices []int) int {
	minPrice := math.MaxInt32
	max := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			max = Max(max, price-minPrice)
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
