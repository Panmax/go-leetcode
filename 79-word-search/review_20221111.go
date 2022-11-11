package main

func exist_20221111(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if word[k] != board[i][j] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		defer func() { visited[i][j] = false }()
		for _, dir := range dirs {
			newI := i + dir.x
			newJ := j + dir.y
			if newI >= 0 && newI < h && newJ >= 0 && newJ < w && !visited[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
