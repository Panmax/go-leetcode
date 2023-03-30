package main

func stoneGameII(piles []int) int {
	n := len(piles)
	s := make([]int, n+1)
	f := make([][]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + piles[i]
		f[i] = make([]int, n+1)
	}

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		// 如果剩余的不足2*m，那么这次就可以拿走 s[n]-s[i] 个
		if 2*m >= n-i {
			return s[n] - s[i]
		}
		// 如果已经有值直接使用
		if f[i][m] > 0 {
			return f[i][m]
		}
		// 否则计算从 f[i][m] 最佳情况，即：
		// 尝试取1..m*2次，每一次可以拿的最佳值为  max(f[i][m],s[n]-s[i]-dfs(i+x, max(x,m)))，也就是我尽量多拿、对方尽量少拿
		for x := 1; x <= m*2; x++ {
			// s[n]-s[i]-dfs(i+x, max(x, m))
			f[i][m] = max(f[i][m], s[n]-s[i]-dfs(i+x, max(x, m)))
		}
		// f[i][m]
		return f[i][m]
	}

	return dfs(0, 1)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
