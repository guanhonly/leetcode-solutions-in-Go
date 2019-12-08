/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
 */
package DynamicProgramming

/**
 * Regular dp solution, dp array is used.
 * The first dimension is day, second is the times that
 * transaction executed, and the third is if a stock is
 * held. But it FAILS since OUT OF MEMORY.
 */
func maxProfitIV_Fail(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	maxProfit := make([][][2]int, n)
	for i := 0; i < len(maxProfit); i++ {
		maxProfit[i] = make([][2]int, k+1)
	}
	INT_MIN := ^int(^uint32(0) >> 1)
	maxProfit[0][0][0] = 0
	maxProfit[0][0][1] = -prices[0]
	for i := 1; i <= k; i++ {
		maxProfit[0][i][0] = INT_MIN
		maxProfit[0][i][1] = INT_MIN
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				maxProfit[i][j][0] = maxProfit[i-1][j][0]
			} else {
				maxProfit[i][j][0] = max(maxProfit[i-1][j-1][1]+prices[i], maxProfit[i-1][j][0])
			}
			maxProfit[i][j][1] = max(maxProfit[i-1][j][0]-prices[i], maxProfit[i-1][j][1])
		}
	}
	res := maxProfit[n-1][0][0]
	for i := 0; i < k; i++ {
		res = max(res, maxProfit[n-1][i][0])
	}
	return res
}

/**
 * A optimized way, which only store the status last time.
 * Time complexity: O(n*k), n is the length of prices.
 * Space complexity: O(n*k)
 */
func maxProfitIV(k int, prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	if k >= n {
		return profits(prices)
	}
	maxProfit := make([][]int, k+1)
	for i := 0; i < len(maxProfit); i++ {
		maxProfit[i] = make([]int, n)
	}
	res := 0
	for i := 1; i <= k; i++ {
		tmpMax := maxProfit[i-1][0] - prices[0] // tmpMax is the status that hold a stock
		for j := 1; j < n; j++ {
			maxProfit[i][j] = max(maxProfit[i][j-1], tmpMax+prices[j])
			tmpMax = max(tmpMax, maxProfit[i-1][j]-prices[j])
			res = max(res, maxProfit[i][j])
		}
	}
	return res
}

func profits(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			res += tmp
		}
	}

	return res
}
