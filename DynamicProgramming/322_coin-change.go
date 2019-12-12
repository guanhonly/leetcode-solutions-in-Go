/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/coin-change/
 */
package DynamicProgramming

/**
 * DP, dp[i] is the fewest size of amount i.
 * Time complexity: O(n*amount), n is the length of coins
 * Space complexity: O(amount)
 */
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
