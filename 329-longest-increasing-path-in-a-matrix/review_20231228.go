package main

func longestIncreasingPath_20231228(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	memory := make([][]int, row)
	for i := range memory {
		memory[i] = make([]int, col)
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memory[r][c] == 0 {
			memory[r][c]++
			for _, dir := range dirs {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < row && newC >= 0 && newC < col && matrix[newR][newC] > matrix[r][c] {
					memory[r][c] = max(memory[r][c], dfs(newR, newC)+1)
				}
			}
		}

		return memory[r][c]
	}

	var ans int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}
