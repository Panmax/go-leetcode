package main

func numIslands_20230608(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= 0 && r < nr && c >= 0 && c < nc && grid[r][c] == '1' {
			grid[r][c] = '0'
			dfs(r, c+1)
			dfs(r, c-1)
			dfs(r+1, c)
			dfs(r-1, c)
		}
	}

	var res int
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				res++
				dfs(r, c)
			}
		}
	}
	return res
}
