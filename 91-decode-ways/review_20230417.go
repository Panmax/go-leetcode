package main

func numDecodings_20230417(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		// s[i-1]
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		// s[i-2]
		if i-2 >= 0 && s[i-2] != '0' && ((s[i-2]-'0')*10+s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
