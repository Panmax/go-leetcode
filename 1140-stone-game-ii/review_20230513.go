package main

func stoneGameII_20230513(piles []int) int {
	n := len(piles)
	s := make([]int, n+1)
	f := make([][]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + piles[i]
		f[i] = make([]int, n+1)
	}
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+2*m >= n {
			return s[n] - s[i]
		}
		if f[i][m] > 0 {
			return f[i][m]
		}
		for x := 1; x <= 2*m; x++ {
			f[i][m] = max(f[i][m], s[n]-s[i]-dfs(i+x, max(x, m)))
		}
		return f[i][m]
	}

	return dfs(0, 1)
}
