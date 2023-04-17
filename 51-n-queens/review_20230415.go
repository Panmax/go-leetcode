package main

func solveNQueens_20230415(n int) [][]string {
	res = make([][]string, 0)
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}
	columns, diagonals1, diagonals2 := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	backtrace_20230415(queens, n, 0, columns, diagonals1, diagonals2)
	return res
}

func backtrace_20230415(queens []int, n int, row int, columns map[int]bool, diagonals1 map[int]bool, diagonals2 map[int]bool) {
	if n == row {
		board := generateBoard_20230415(queens, n)
		res = append(res, board)
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
		columns[i] = true
		diagonals1[diagonal1] = true
		diagonals2[diagonal2] = true
		queens[row] = i
		backtrace_20230415(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard_20230415(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := range row {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board[i] = string(row)
	}
	return board
}
