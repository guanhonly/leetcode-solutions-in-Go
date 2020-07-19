package DFSAndBFS

var (
    UNCOLORED, RED, GREEN = 0, 1, 2
    color []int
    valid bool
)

func isBipartite(graph [][]int) bool {
    n := len(graph)
    valid = true
    color = make([]int, n)
    for i := 0; i < n && valid; i++ {
        if color[i] == UNCOLORED {
            dfs2(i, RED, graph)
        }
    }
    return valid
}

func dfs2(node, c int, graph [][]int) {
    color[node] = c
    cNei := RED
    if c == RED {
        cNei = GREEN
    }
    for _, neighbor := range graph[node] {
        if color[neighbor] == UNCOLORED {
            dfs2(neighbor, cNei, graph)
            if !valid {
                return
            }
        } else if color[neighbor] != cNei {
            valid = false
            return
        }
    }
}