package main

import "math"

func maxProfit_20230316(prices []int) int {
	minPrice := math.MaxInt32
	res := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			res = Max(res, price-minPrice)
		}
	}
	return res
}
