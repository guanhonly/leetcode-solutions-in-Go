/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/uncrossed-lines/
 */

package DynamicProgramming

// DP.
func maxUncrossedLines(A []int, B []int) int {
	m := len(A)
	n := len(B)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[m][n]
}
