package DFSAndBFS

var dx2 = []int{0, -1, 1, 0, -1, -1, 1, 1}
var dy2 = []int{-1, 0, 0, 1, -1, 1, -1, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x := click[0]
	y := click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	if board[x][y] == 'B' || board[x][y] == 'X' || (board[x][y]-'0' >= 0 && board[x][y]-'0' <= 9) {
		return board
	}
	dfs6(board, x, y)
	return board
}

func dfs6(board [][]byte, x, y int) {
	if !isValid2(x, y, len(board), len(board[0])) || board[x][y] != 'E' {
		return
	}
	if board[x][y] == 'E' {
		board[x][y] = 'B'
		minesNumber := nearbyMinesNumber(board, x, y)
		if minesNumber > 0 {
			board[x][y] = '0' + minesNumber
			// if current place is number, return
			return
		}
	}
	for i := 0; i < 8; i++ {
		dfs6(board, x+dx2[i], y+dy2[i])
	}
}

func nearbyMinesNumber(board [][]byte, x, y int) byte {
	res := 0
	for i := 0; i < 8; i++ {
		newX, newY := x+dx2[i], y+dy2[i]
		if !isValid2(newX, newY, len(board), len(board[0])) {
			continue
		}
		if board[newX][newY] == 'M' {
			res++
		}
	}
	return byte(res)
}

func isValid2(x, y, m, n int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	return true
}
