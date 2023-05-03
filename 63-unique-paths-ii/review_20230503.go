package main

func uniquePathsWithObstacles_20230503(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			// dp
			dp[i][0] = 1
		} else {
			break
		}
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			// dp
			dp[0][i] = 1
		} else {
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 判断1
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
