package TwoPointers

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	const INT_MAX = int(^uint(0) >> 1)
	start, end := 0, 0
	ans := INT_MAX
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == INT_MAX {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
