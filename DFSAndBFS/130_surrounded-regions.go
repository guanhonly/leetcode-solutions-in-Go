package DFSAndBFS

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])

	for i := 0; i < m; i++ {
		dfs4(i, 0, board)
		dfs4(i, n-1, board)
	}

	for i := 1; i < n-1; i++ {
		dfs4(0, i, board)
		dfs4(m-1, i, board)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs4(x, y int, board [][]byte) {
	if x >= len(board) || x < 0 || y >= len(board[0]) || y < 0 || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	for i := 0; i < 4; i++ {
		dfs4(x+dx[i], y+dy[i], board)
	}
}
