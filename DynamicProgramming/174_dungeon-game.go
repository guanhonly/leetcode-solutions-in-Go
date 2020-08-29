package DynamicProgramming

// DP, the path from the top-left to the bottom-right is not
// appropriate because two metric need to consider, one is the
// sum to the i,j, another is the minimum health point on the start
// position. So this DP is from the bottom-right to the top-left,
// which only records the minimum point from the start position.
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	const INT_MAX = int(^uint(0) >> 1)
	n := len(dungeon[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = INT_MAX
		}
	}
	dp[m][n-1], dp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			minn := min(dp[i][j+1], dp[i+1][j])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}
