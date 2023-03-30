package main

import "math"

func review_20230330(coins []int, amount int) int {
	rowNum, colNum := len(coins)+1, amount+1
	dp := make([][]int, rowNum)
	for i := range dp {
		dp[i] = make([]int, colNum)
	}
	max := math.MaxInt32
	for i := 1; i < colNum; i++ {
		dp[0][i] = max
	}

	for row := 1; row < rowNum; row++ {
		coin := coins[row-1]
		for col := 1; col < colNum; col++ {
			if coin > col {
				dp[row][col] = dp[row-1][col]
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row][col-coin]+1)
			}
		}
	}
	res := dp[rowNum-1][colNum-1]
	if res == max {
		return -1
	}
	return res
}
