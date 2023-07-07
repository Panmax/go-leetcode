package main

func solve_20230707(board [][]byte) {
	row, col := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= 0 && r < row && c >= 0 && c < col && board[r][c] == 'O' {
			board[r][c] = 'A'
			dfs(r-1, c)
			dfs(r+1, c)
			dfs(r, c-1)
			dfs(r, c+1)
		}
	}
	for i := 0; i < row; i++ {
		dfs(i, 0)
		dfs(i, col-1)
	}
	for i := 0; i < col; i++ {
		dfs(0, i)
		dfs(row-1, i)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == 'A' {
				board[r][c] = 'O'
			} else if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}
}
