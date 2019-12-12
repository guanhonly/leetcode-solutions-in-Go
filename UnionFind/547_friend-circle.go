package UnionFind

/**
 * DFS. This question is similar to number-of-islands.
 * Time complexity: O(n^2), n is the length of M
 * Space complexity: O(n)
 */
func findCircleNum(M [][]int) int {
    m := len(M)
    if m == 0 {
        return 0
    }
    if len(M[0]) !=m  {
        return 0
    }
    visited := make([]bool, m)
    count := 0
    for i:=0; i<m; i++ {
        if !visited[i] {
            dfs2(M, visited, i)
            count++
        }
    }
    return count
}

func dfs2(M [][]int, visited []bool, row int) {
    for i:=0; i<len(M); i++ {
        if M[row][i] == 1 && !visited[i] {
            visited[i] = true
            dfs2(M, visited, i)
        }
    }
}

/**
 * UnionFind.
 * Time complexity: O(n^3)
 * Space complexity: O(n)
 */
func findCircleNum2(M [][]int) int {
    m := len(M)
    if m == 0 {
        return 0
    }
    uf := NewUnionFind2(M)
    for i:=0; i<m; i++ {
        for j:=0; j<m; j++ {
            if M[i][j] == 1 {
                uf.Union(i, j)
            }
        }
    }
    return uf.GetCount()
}
type unionFind2 struct {
    roots []int
    ranks []int
    count int
}
func NewUnionFind2(M [][]int) *unionFind2{
    uf := unionFind2{}
    m := len(M)
    uf.count =m

    uf.roots = make([]int, m)
    uf.ranks = make([]int, m)
    for i:=0; i<m; i++ {
        uf.roots[i] = i
        uf.ranks[i] = 0
    }
    return &uf
}

func (s *unionFind2) find(i int) int {
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

func (s *unionFind2) Union(x, y int) {
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

func (s *unionFind2) GetCount() int {
    return s.count
}