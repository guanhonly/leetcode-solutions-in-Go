/**
 * Difficulty: Medium
 * Question link: https://leetcode-cn.com/problems/generate-parentheses/
 */
package DFSAndBFS

func generateParenthesis(n int) []string {
	var res []string
	backTrack(&res, "", 0, 0, n)
	return res
}

func backTrack(res *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*res = append(*res, current)
		return
	}
	if open < max {
		backTrack(res, current+"(", open+1, close, max)
	}
	if close < open {
		backTrack(res, current+")", open, close+1, max)
	}
}

func generateParenthesis2(n int) []string {
	var res []string
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range generateParenthesis2(i) {
				for _, right := range generateParenthesis2(n - 1 - i) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
