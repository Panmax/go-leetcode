package main

func isMatch_20221227(s string, p string) bool {
	n1, n2 := len(s), len(p)
	dp := make([][]bool, n1+1)
	for i := range dp {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true
	for i := 0; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			if p[j-1] != '*' {
				if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if i > 0 && j > 1 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i-1][j]
				}
				if j > 1 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
			}
		}
	}
	return dp[n1][n2]
}
