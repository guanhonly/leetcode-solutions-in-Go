/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 */
package DynamicProgramming

/**
 * Regular dp.
 * Time complexity: O(n), n is the length of prices
 * Space complexity: O(n)
 */
func maxProfitWithTransaction(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := make([][2]int, n)
	maxProfit[0][0], maxProfit[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		maxProfit[i][0] = max(maxProfit[i-1][1]+prices[i]-fee, maxProfit[i-1][0])
		maxProfit[i][1] = max(maxProfit[i-1][0]-prices[i], maxProfit[i-1][1])
	}
	return maxProfit[n-1][0]
}

/**
 * Optimized way, only two variables are used.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func maxProfitWithTransaction2(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	hold, release := -prices[0], 0
	for i := 1; i < n; i++ {
		release = max(release, hold+prices[i]-fee)
		hold = max(hold, release-prices[i])
	}
	return release
}
