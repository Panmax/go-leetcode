package main

func solveNQueens_20231112(n int) [][]string {
	res = make([][]string, 0)
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make(map[int]bool)
	diagonals1 := make(map[int]bool)
	diagonals2 := make(map[int]bool)
	backtrace_20231112(queens, n, 0, columns, diagonals1, diagonals2)

	return res
}

func backtrace_20231112(queens []int, n int, row int, columns map[int]bool, diagonals1 map[int]bool, diagonals2 map[int]bool) {
	if n == row {
		boards := generateBoard_20231112(queens, n)
		res = append(res, boards)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1] = true
		diagonals2[diagonal2] = true
		backtrace_20231112(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard_20231112(queens []int, n int) []string {
	var boards []string
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		boards = append(boards, string(row))
	}
	return boards
}
