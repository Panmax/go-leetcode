package main

func maximumWealth_20230727(accounts [][]int) int {
	max := 0
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
