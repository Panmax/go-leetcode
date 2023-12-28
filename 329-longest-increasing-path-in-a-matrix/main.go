package main

var dirs = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	row, column := len(matrix), len(matrix[0])
	memory := make([][]int, row)
	for i := range memory {
		memory[i] = make([]int, column)
	}
	var ans int
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memory[r][c] > 0 {
			return memory[r][c]
		}
		memory[r][c] = 1
		for _, dir := range dirs {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < column && matrix[newR][newC] > matrix[r][c] {
				memory[r][c] = max(memory[r][c], dfs(newR, newC)+1)
			}
		}
		return memory[r][c]
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			ans = max(ans, dfs(i, j))
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
