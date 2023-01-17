package main

import "math"

func review_20230113(coins []int, amount int) int {
	max := math.MaxInt32
	// +1
	rowSize, colSize := len(coins)+1, amount+1
	dp := make([][]int, rowSize)
	for i := range dp {
		dp[i] = make([]int, colSize)
	}
	for i := 1; i < colSize; i++ {
		dp[0][i] = max
	}
	for row := 1; row < rowSize; row++ {
		coin := coins[row-1]
		for col := 1; col < colSize; col++ {
			if coin > col {
				dp[row][col] = dp[row-1][col]
			} else {
				dp[row][col] = min(dp[row-1][col], dp[row][col-coin]+1)
			}
		}
	}
	// res 判断
	res := dp[rowSize-1][colSize-1]
	if res == max {
		res = -1
	}
	return res
}
