package main

func longestIncreasingPath_20240417(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	memory := make([][]int, m)
	for i := range memory {
		memory[i] = make([]int, n)
	}
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if memory[r][c] == 0 {
			memory[r][c]++
			for _, dir := range dirs {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < m && newC >= 0 && newC < n && matrix[newR][newC] > matrix[r][c] {
					memory[r][c] = max(memory[r][c], dfs(newR, newC)+1)
				}
			}
		}
		return memory[r][c]
	}

	var ans int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}

	return ans
}
