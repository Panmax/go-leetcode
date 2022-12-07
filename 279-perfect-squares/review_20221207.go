package main

func numSquares_20221207(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			// 需要+1
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
