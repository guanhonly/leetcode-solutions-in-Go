/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/triangle/
 */
package DynamicProgramming

/**
 * The common solution using dp. the dp[2][2] stores
 * the min value and max value of prev and current, which
 * means that dp[y][0] is the max value of prev, dp[y][1]
 * is the min value of current, dp[x][0] is the max value of
 * current, and dp[x][1] is the min value of current.
 * Time complexity: O(n), n is the length of nums
 * Space complexity: O(1), since only const variables are used.
 */
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var dp [2][2]int
	var res int
	dp[0][0], dp[0][1], res = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x, y := i&1, (i-1)&1
		dp[x][0] = max(dp[y][0]*nums[i], dp[y][1]*nums[i], nums[i])
		dp[x][1] = min(dp[y][0]*nums[i], dp[y][1]*nums[i], nums[i])
		res = max(res, dp[x][0])
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

/**
 * A more concise way, which use only two variables to replace
 * the dp array.
 */
func maxProduct2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	currMax := res
	currMin := res
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax
		}
		currMax = max(currMax*nums[i], nums[i])
		currMin = max(currMin*nums[i], nums[i])
		res = max(res, currMax)
	}
	return res
}
