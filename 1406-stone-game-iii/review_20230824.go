package main

import "math"

func stoneGameIII_20230824(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)
	var sum int

	for i := n - 1; i >= 0; i-- {
		dp[i] = math.MinInt32
		sum += stoneValue[i]
		for j := i; j < i+3 && j < n; j++ {
			dp[i] = max(dp[i], sum-dp[j+1])
		}
	}

	alias := dp[0]
	bob := sum - dp[0]
	if alias > bob {
		return "Alice"
	} else if alias < bob {
		return "Bob"
	}
	return "Tie"
}
