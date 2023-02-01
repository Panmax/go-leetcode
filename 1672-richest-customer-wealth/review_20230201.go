package main

func maximumWealth_20230201(accounts [][]int) int {
	max := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for _, val := range accounts[i] {
			sum += val
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
