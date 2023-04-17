package main

func solve_20230417(board [][]byte) {
	row, col := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= row || j < 0 || j >= col || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'A'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < row; i++ {
		dfs(i, 0)
		// 右边界
		dfs(i, col-1)
	}
	for i := 0; i < col; i++ {
		dfs(0, i)
		// 下边界
		dfs(row-1, i)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
