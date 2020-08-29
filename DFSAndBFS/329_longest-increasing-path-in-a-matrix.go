package DFSAndBFS

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, -1, 0, 1}

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	mem := make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs3(matrix, mem, i, j))
		}
	}
	return res
}

func dfs3(matrix, mem [][]int, startX, startY int) int {
	if mem[startX][startY] != 0 {
		return mem[startX][startY]
	}
	mem[startX][startY]++
	for i := 0; i < 4; i++ {
		nextX := startX + dx[i]
		nextY := startY + dy[i]
		if nextX >= 0 && nextX < len(matrix) && nextY >= 0 && nextY < len(matrix[0]) &&
			matrix[nextX][nextY] > matrix[startX][startY] {
			mem[startX][startY] = max(mem[startX][startY], dfs3(matrix, mem, nextX, nextY)+1)
		}
	}
	return mem[startX][startY]
}
