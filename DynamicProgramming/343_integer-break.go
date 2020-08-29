package DynamicProgramming

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		var currMax int
		for j := 1; j < i; j++ {
			currMax = max(currMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = currMax
	}
	return dp[n]
}
