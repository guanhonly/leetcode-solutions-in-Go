/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
 */
package DynamicProgramming

/**
 * DP method, the dp array is used.
 * For the array, the first dimension is the day, the second is if a stock
 * is engaged, the third is the times of transactions.
 * Time complexity: O(n), n is the length of prices
 * Space complexity: O(n)
 */
func maxProfitIII(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	// it's not the appropriate way, but Go dose not handle overflow as Python does.
	INT_MIN := ^int(^uint32(0) >> 1)
	maxProfit := make([][2][3]int, n)
	maxProfit[0][0][0], maxProfit[0][1][0] = 0, -prices[0]
	maxProfit[0][0][1], maxProfit[0][1][1] = INT_MIN, INT_MIN
	maxProfit[0][1][2], maxProfit[0][1][2] = INT_MIN, INT_MIN
	for i := 1; i < n; i++ {
		maxProfit[i][0][0] = maxProfit[i-1][0][0]
		maxProfit[i][0][1] = max(maxProfit[i-1][1][0]+prices[i], maxProfit[i-1][0][1])
		maxProfit[i][0][2] = max(maxProfit[i-1][1][1]+prices[i], maxProfit[i-1][0][2])

		maxProfit[i][1][0] = max(maxProfit[i-1][1][0], maxProfit[i-1][0][0]-prices[i])
		maxProfit[i][1][1] = max(maxProfit[i-1][1][1], maxProfit[i-1][0][1]-prices[i])
	}
	return max(maxProfit[n-1][0][2], maxProfit[n-1][0][1], maxProfit[n-1][0][0])
}

/**
 * A more concise way, without using dp array but 4 variables.
 * hold1 is the first time of buying stock and hold2 is the second time
 * of buying stock. release1 is the first time of selling stock and release2
 * is the second time of selling stock.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func maxProfitIII2(prices []int) int {
	INT_MIN := ^int(^uint32(0) >> 1)
	hold1, hold2 := INT_MIN, INT_MIN
	release1, release2 := 0, 0
	for _, i := range prices {
		release2 = max(release2, hold2+i)
		hold2 = max(hold2, release1-i)
		release1 = max(release1, hold1+i)
		hold1 = max(hold1, -i)
	}
	return release2
}
