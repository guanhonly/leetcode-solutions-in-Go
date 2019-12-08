/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
 */
package DynamicProgramming

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	lowestPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			lowestPrice = min(lowestPrice, prices[i])
		} else {
			profit = max(profit, prices[i]-lowestPrice)
		}
	}
	return profit
}
