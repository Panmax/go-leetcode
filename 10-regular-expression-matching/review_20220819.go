package main

func isMatch_20220819(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	// 注意这里初始化
	dp[0][0] = true

	for i := 0; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] != '*' {
				if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				if i > 0 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i-1][j]
				}

				// 注意下边不用判断 i > 0
				if j >= 2 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
			}
		}
	}

	return dp[m][n]
}