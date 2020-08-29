package BinarySearch

/* Prefix sum, if sums[bound] - sums[i] >= s, means the
 * numbers between the bound and i is the subarray we want.
 */
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	const INT_MAX = int(^uint(0) >> 1)
	ans := INT_MAX
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := lowerBound(sums, target)
		if bound < 0 {
			continue
		}
		if bound <= n {
			ans = min(ans, bound-i+1)
		}
	}
	if ans == INT_MAX {
		return 0
	}
	return ans
}

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] >= target {
		return left
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
