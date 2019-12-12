/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/number-of-islands/
 */
package UnionFind

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}
var m, n int
var visited [][]bool
var gridG [][]byte

/**
 * DFS, the most efficient way.
 * Time complexity: O(m*n), m is the rows of grid with n is the columns
 * Space complexity: O(m*n)
 */
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	gridG = grid
	m, n = len(grid), len(grid[0])
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count += dfs(i, j)
		}
	}
	return count
}

func dfs(x, y int) int {
	if !isValid(x, y) {
		return 0
	}
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		dfs(x+dx[i], y+dy[i])
	}
	return 1
}

func isValid(x, y int) bool {
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	if gridG[x][y] == '0' || visited[x][y] {
		return false
	}
	return true
}

/**
 * BFS.
 * Time complexity: O(m*n)
 * Space complexity: O(m*n)
 */
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	gridG = grid
	m, n = len(grid), len(grid[0])
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count += bfs(i, j)
		}
	}
	return count
}

func bfs(x, y int) int {
	if !isValid(x, y) {
		return 0
	}
	visited[x][y] = true
	var queue [][]int
	queue = append(queue, []int{x, y})
	for len(queue) > 0 {
		curX, curY := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			newX, newY := curX+dx[i], curY+dy[i]
			if isValid(newX, newY) {
				visited[newX][newY] = true
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	return 1
}

/**
 * UnionFind way, not effective enough.
 * Time complexity: O(m*n)
 * Space complexity: O(m*n)
 */
func numIslands3(grid [][]byte) int {
	m = len(grid)
	if m == 0 {
		return 0
	}
	n = len(grid[0])
	if n == 0 {
		return 0
	}
	gridG = grid
	uf := NewUnionFind(grid)
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				visited[i][j] = true
				for k := 0; k < 4; k++ {
					nextX := i + dx[k]
					nextY := j + dy[k]
					if isValid(nextX, nextY) {
						uf.Union(i*n+j, nextX*n+nextY)
					}
				}
			}
		}
	}
	return uf.GetCount()
}

type unionFind struct {
	roots []int
	ranks []int
	count int
}

func NewUnionFind(grid [][]byte) *unionFind {
	uf := unionFind{}
	uf.count = 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.roots = append(uf.roots, i*n+j)
				uf.count++
			} else {
				uf.roots = append(uf.roots, -1)
			}
			uf.ranks = append(uf.ranks, 0)
		}
	}
	return &uf
}

func (s *unionFind) find(i int) int {
	root := i
	// find root
	for root != s.roots[root] {
		root = s.roots[root]
	}
	// compress path
	for i != s.roots[i] {
		i, s.roots[i] = s.roots[i], root
	}
	return root
}

func (s *unionFind) Union(x, y int) {
	xRoot := s.find(x)
	yRoot := s.find(y)
	if xRoot != yRoot {
		if s.ranks[xRoot] > s.ranks[yRoot] {
			s.roots[yRoot] = xRoot
		} else if s.ranks[xRoot] < s.ranks[yRoot] {
			s.roots[xRoot] = yRoot
		} else {
			s.roots[yRoot] = xRoot
			s.ranks[xRoot]++
		}
		s.count--
	}
}

func (s *unionFind) GetCount() int {
	return s.count
}
