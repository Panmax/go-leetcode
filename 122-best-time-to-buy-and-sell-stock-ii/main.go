package main

// dp[i][0]表示第 i 天交易完后手里没有股票的最大利润
// dp[i][1]表示第 i 天交易完后手里持有股票的最大利润
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
