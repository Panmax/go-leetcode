package main

func numDecodings_20240407(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i < len(s)+1; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0')) <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
