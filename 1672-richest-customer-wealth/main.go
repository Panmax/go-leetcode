package main

import "math"

func maximumWealth(accounts [][]int) int {
	max := math.MinInt32
	for _, account := range accounts {
		sum := 0
		for _, num := range account {
			sum += num
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
