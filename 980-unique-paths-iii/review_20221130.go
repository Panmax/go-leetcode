package main

func uniquePathsIII_20221130(grid [][]int) int {
	w, h := len(grid), len(grid[0])

	var todo, total int
	var sr, sc, tr, tc int
	for r, row := range grid {
		for c, val := range row {
			if val != -1 {
				todo++
			}
			if val == 1 {
				sr = r
				sc = c
			}
			if val == 2 {
				tr = r
				tc = c
			}
		}
	}

	var dfs func(r, c, todo int)
	dfs = func(r, c, todo int) {
		// 记得要--
		todo--
		if todo < 0 {
			return
		}
		if r == tr && c == tc {
			if todo == 0 {
				total++
			}
			return
		}
		grid[r][c] = -1
		for _, dir := range dirs {
			newR := r + dir.r
			newC := c + dir.c
			if newR >= 0 && newR < w && newC >= 0 && newC < h && grid[newR][newC]%2 == 0 {
				dfs(newR, newC, todo)
			}
		}
		grid[r][c] = 0
	}

	dfs(sr, sc, todo)
	return total
}
