package main

func uniquePathsIII_20230130(grid [][]int) int {
	w, h := len(grid), len(grid[0])
	var sr, sc, tr, tc int
	var todo, total int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != -1 {
				todo++
			}
			if grid[i][j] == 1 {
				sr = i
				sc = j
			}
			if grid[i][j] == 2 {
				tr = i
				tc = j
			}
		}
	}
	// 最好用 r, c
	var dfs func(i, j, todo int)
	dfs = func(i, j, todo int) {
		todo--
		if todo < 0 {
			return
		}
		if todo == 0 {
			if i == tr && j == tc {
				total++
			}
			return
		}
		grid[i][j] = -1
		for _, dir := range dirs {
			newI := i + dir.r
			newJ := j + dir.c
			if newI >= 0 && newI < w && newJ >= 0 && newJ < h && grid[newI][newJ] != -1 {
				dfs(newI, newJ, todo)
			}
		}
		grid[i][j] = 0
	}
	dfs(sr, sc, todo)

	return total
}
