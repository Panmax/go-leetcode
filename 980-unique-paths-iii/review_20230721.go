package main

func uniquePathsIII_20230721(grid [][]int) int {
	w, h := len(grid), len(grid[0])
	var todo, total int
	var sr, sc int
	var tr, tc int
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if grid[i][j] != -1 {
				todo++
			}
			// if
			if grid[i][j] == 1 {
				sr, sc = i, j
			}
			// if
			if grid[i][j] == 2 {
				tr, tc = i, j
			}
		}
	}

	var dfs func(r, c, todo int)
	dfs = func(r int, c int, todo int) {
		todo--
		if todo < 0 {
			return
		}
		if todo == 0 {
			if r == tr && c == tc {
				total++
			}
			return
		}
		grid[r][c] = -1
		for _, dir := range dirs {
			newR := r + dir.r
			newC := c + dir.c
			if newR >= 0 && newR < w && newC >= 0 && newC < h && grid[newR][newC] != -1 {
				dfs(newR, newC, todo)
			}
		}
		grid[r][c] = 0
	}

	dfs(sr, sc, todo)

	return total
}
