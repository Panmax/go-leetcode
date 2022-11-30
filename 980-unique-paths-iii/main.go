package main

import "fmt"

type pair struct {
	r, c int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func uniquePathsIII(grid [][]int) int {
	w, h := len(grid), len(grid[0])

	var todo, total int
	var sr, sc int
	var tr, tc int

	for r, row := range grid {
		for c, val := range row {
			if val != -1 {
				todo++
			}
			if val == 1 {
				sr, sc = r, c
			}
			if val == 2 {
				tr, tc = r, c
			}
		}
	}

	var dfs func(r, c, todo int)
	dfs = func(r, c, todo int) {
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

func main() {
	fmt.Println(uniquePathsIII([][]int{{1, 2}}))
}
