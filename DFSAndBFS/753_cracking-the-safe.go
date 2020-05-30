/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/cracking-the-safe/
 */
package DFSAndBFS

import (
	"strconv"
	"strings"
)

/**
 * Hierholzer algorithm, to  finding an Euler path (a path
 * visiting every edge exactly once) on the following graph:
 * there are k^(n-1) nodes with each node having k edges.
 * Time complexity: O(n*k^n)
 * Space complexity: O(n*k^n)
 */
func crackSafe(n int, k int) string {
	if n == 1 && k == 1 {
		return "0"
	}
	visited := make(map[string]bool)
	res := ""
	var node string
	for i := 0; i < n-1; i++ {
		node += "0"
	}
	dfs(node, &res, k, visited)
	res += node
	return res
}

func dfs(node string, res *string, k int, visited map[string]bool) {
	for i := 0; i < k; i++ {
		newNode := node + strconv.Itoa(i)
		if !visited[newNode] {
			visited[newNode] = true
			dfs(newNode[1:], res, k, visited)
			*res = *res + strconv.Itoa(i)
		}
	}
}

/**
 * Another way, looks like more brute but more efficient.
 */
func crackSafe2(n int, k int) string {
	res := strings.Repeat("0", n)
	visited := make(map[string]bool)
	visited[res] = true
	for i := 0; i < pow(k, n); i++ {
		prev := res[len(res)-n+1:]
		for j := k - 1; j >= 0; j-- {
			curr := prev + strconv.Itoa(j)
			if !visited[curr] {
				visited[curr] = true
				res += strconv.Itoa(j)
				break
			}
		}
	}
	return res
}

func pow(x, n int) int {
	res := 1
	tmp := x
	for i := n; i >= 1; i /= 2 {
		if i&1 == 1 {
			res *= tmp
		}
		tmp *= tmp
	}
	return res
}
