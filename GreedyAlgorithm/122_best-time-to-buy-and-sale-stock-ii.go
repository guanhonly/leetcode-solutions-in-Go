/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
 */
package GreedyAlgorithm

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
