package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// init
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else { // break
			break
		}
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else { // break
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 判断
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1]
}
