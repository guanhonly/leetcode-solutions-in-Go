/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/climbing-stairs/
 */
package DynamicProgramming

/**
 * An intuitive dp.
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/**
 * Optimized dp, the useful elements in dp array are only
 * the prev and prev's prev.
 */
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	before := 1
	allWays := 2
	for i := 2; i < n; i++ {
		allWays, before = allWays+before, allWays
	}
	return allWays
}
