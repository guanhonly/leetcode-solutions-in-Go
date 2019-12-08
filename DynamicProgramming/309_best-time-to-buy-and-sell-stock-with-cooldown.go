/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
 */
package DynamicProgramming

/*
 * The first dimension is day, the second is the status that if
 * transaction is allowed, the third is the status that if a stock
 * is held.
 * Time complexity: O(n), n is the length of prices.
 * Space complexity: O(n)
 */
func maxProfitWithCooldown(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	INT_MIN := ^int(^uint32(0) >> 1)
	maxProfit := make([][2][2]int, n)
	// 0 is the transaction allowed, 1 is not allowed.
	maxProfit[0][0][0] = 0
	maxProfit[0][0][1] = -prices[0]
	maxProfit[0][1][0] = INT_MIN
	maxProfit[0][1][1] = INT_MIN
	for i := 1; i < n; i++ {
		maxProfit[i][0][0] = max(maxProfit[i-1][1][0], maxProfit[i-1][0][0])
		maxProfit[i][0][1] = max(maxProfit[i-1][0][1], maxProfit[i-1][0][0]-prices[i])

		maxProfit[i][1][0] = maxProfit[i-1][0][1] + prices[i]
	}
	return max(maxProfit[n-1][0][0], maxProfit[n-1][1][0])
}
