package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 注意 x,y 边界条件
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 0; i < m; i++ {
		dfs(0, i)
		dfs(n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// A -> O ; O -> X
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
