package main

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns, diagonals1, diagonals2 := make(map[int]bool, 0), make(map[int]bool, 0), make(map[int]bool, 0)
	backtrace(queens, n, 0, columns, diagonals1, diagonals2)
	return res
}

func backtrace(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if n == row {
		board := generateBoard(queens, n)
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
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1] = true
		diagonals2[diagonal2] = true
		backtrace(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board[i] = string(row)
	}
	return board
}
