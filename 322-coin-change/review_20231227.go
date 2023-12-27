package main

import "math"

func review_20231227(coins []int, amount int) int {
	rowNum, colNum := len(coins)+1, amount+1
	dp := make([][]int, rowNum)
	for i := range dp {
		dp[i] = make([]int, colNum)
	}
	for i := 0; i < colNum; i++ {
		dp[0][i] = math.MaxInt32
	}

	for row := 1; row < rowNum; row++ {
		for col := 1; col < colNum; col++ {
			coin := coins[row-1]
			if coin > col {
				dp[row][col] = dp[row-1][col]
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row][col-coin]+1)
			}
		}
	}

	res := dp[rowNum-1][colNum-1]
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
