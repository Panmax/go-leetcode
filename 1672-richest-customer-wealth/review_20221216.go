package main

func maximumWealth_20221216(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		var sum int
		for _, val := range account {
			sum += val
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
