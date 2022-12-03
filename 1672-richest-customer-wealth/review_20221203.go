package main

import "math"

func maximumWealth_20221203(accounts [][]int) int {
	max := math.MinInt32
	for _, account := range accounts {
		sum := 0
		for _, val := range account {
			sum += val
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
