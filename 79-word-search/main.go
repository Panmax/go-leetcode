package main

type pair struct {
	x, y int
}

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		defer func() { visited[i][j] = false }()
		for _, dir := range dirs {
			newI := dir.x + i
			newJ := dir.y + j
			if newI >= 0 && newI < h && newJ >= 0 && newJ < w && !visited[newI][newJ] {
				// if
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

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	println(exist(board, "ABCCED"))

}
