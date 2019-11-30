/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/word-search-ii/
 */

package Trie

import (
    "strings"
)

/**
 * The original idea: only using dfs. But timeout.
 */
func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 {
        return nil
    }
    if len(board[0]) == 0 {
        return nil
    }
    m := len(board)
    n := len(board[0])
    visited := make([][]bool, m)
    for i := 0; i<m; i++ {
        line := make([]bool, n)
        for j:=0; j<n; j++ {
            line[j] = false
        }
        visited[i] = line
    }
    set := make(map[string]bool)
    dx := [4]int{1, 0, -1, 0}
    dy := [4]int{0, 1, 0, -1}
    for _, word := range words {
        for i := 0; i < m; i++ {
            for j := 0; j < n; j++ {
                dfs(i, j, "", word, board, visited, set, dx, dy)
            }
        }
    }
    keys := make([]string, len(set))
    i := 0
    for k := range set {
        keys[i] = k
        i++
    }
    return keys
}

func dfs(x, y int, curStr, word string, board [][]byte, visited [][]bool, res map[string]bool, dx, dy [4]int)  {
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] {
        return
    }
    curStr = curStr + string(board[x][y])
    if !strings.Contains(word, curStr) {
        return
    }
    if strings.Compare(curStr, word) == 0 {
        res[curStr] = true
    }
    visited[x][y] = true
    for i:=0 ;i<4; i++ {
        newX := x + dx[i]
        newY := y + dy[i]
        dfs(newX, newY, curStr, word, board, visited, res, dx, dy)
    }
    visited[x][y] = false
}


/**
 * Using trie to prune, the struct of trie is in the file 208_implement-trie-prefix-tree.go.
 */
func findWords2(board [][]byte, words []string) []string {
    if len(board) == 0 {
        return nil
    }
    if len(board[0]) == 0 {
        return nil
    }
    m := len(board)
    n := len(board[0])
    visited := make([][]bool, m)
    for i := 0; i<m; i++ {
        line := make([]bool, n)
        for j:=0; j<n; j++ {
            line[j] = false
        }
        visited[i] = line
    }
    trie := Constructor()
    for _, word := range words {
        trie.Insert(word)
    }
    set := make(map[string]bool)
    dx := [4]int{1, 0, -1, 0}
    dy := [4]int{0, 1, 0, -1}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs2(i, j, "", visited, board, trie, set, dx, dy)
        }
    }
    res := make([]string, len(set))
    i := 0
    for k := range set {
        res[i] = k
        i++
    }
    return res
}

func dfs2(x, y int, curStr string, visited [][]bool, board [][]byte, trie Trie, res map[string]bool, dx, dy [4]int) {
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] {
        return
    }
    curStr = curStr + string(board[x][y])
    if !trie.StartsWith(curStr) {
        return
    }
    if trie.Search(curStr) {
        res[curStr] = true
    }
    visited[x][y] = true
    for i:=0 ;i<4; i++ {
        newX := x + dx[i]
        newY := y + dy[i]
        dfs2(newX, newY, curStr, visited, board, trie, res, dx, dy)
    }
    visited[x][y] = false
}

/**
 * It's the more efficient way I found in Accepted Solutions Runtime Distribution,
 * the idea are similar as the way above but used a more light struct of trie.
 */
func findWords3(board [][]byte, words []string) []string {
    var results []string

    m := len(board)
    if m == 0 {
        return results
    }

    n := len(board[0])
    if n == 0 {
        return results
    }

    trie := buildTrie(words)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs3(board, i, j, trie, &results)
        }
    }

    return results
}

func dfs3(board [][]byte, i int, j int, trie *TrieNode2, results *[]string) {
    c := board[i][j]
    if c == '#' || trie.next[int(c-'a')] == nil {
        return
    }

    trie = trie.next[int(c-'a')]
    if len(trie.word) > 0 {
        // Found one
        *results = append(*results, trie.word)
        trie.word = ""
    }

    board[i][j] = '#'
    if i > 0 {
        dfs3(board, i-1, j, trie, results)
    }

    if i < len(board)-1 {
        dfs3(board, i+1, j, trie, results)
    }

    if j > 0 {
        dfs3(board, i, j-1, trie, results)
    }

    if j < len(board[0])-1 {
        dfs3(board, i, j+1, trie, results)
    }

    board[i][j] = c
}

type TrieNode2 struct {
    next [26]*TrieNode2
    word string
}

func buildTrie(words []string) *TrieNode2 {
    root := new(TrieNode2)
    for _, word := range words {
        cur := root
        for _, c := range word {
            idx := int(c-'a')
            if cur.next[idx] == nil {
                cur.next[idx] = new(TrieNode2)
            }
            cur = cur.next[idx]
        }
        cur.word = word
    }

    return root
}