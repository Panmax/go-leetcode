package main

func numIslands_20231101(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= 0 && r < nr && c >= 0 && c < nc && grid[r][c] == '1' {
			grid[r][c] = 0
			dfs(r+1, c)
			dfs(r-1, c)
			dfs(r, c+1)
			dfs(r, c-1)
		}
	}
	var res int
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
