package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] != '*' {
				if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				// 判断 i 大于0， p[j-2] 是 .
				if j >= 2 && i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i-1][j]
				}
				if j >= 2 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}
