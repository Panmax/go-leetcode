package main

import "fmt"

func numIslands(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])
	var res int

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= 0 && i < nr && j >= 0 && j < nc && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}

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

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}

	fmt.Println(numIslands(grid))
}
