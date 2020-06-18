/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/longest-increasing-subsequence/
 */
package DynamicProgramming

/**
 * DP.
 * Time complexity: O(n^2)
 * Space complexity: O(n)
 */
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	LIS := make([]int, n)
	for i := 0; i < n; i++ {
		LIS[i] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				LIS[i] = max(LIS[i], LIS[j]+1)
			}
		}
		maxLen = max(maxLen, LIS[i])
	}
	return maxLen
}

/**
 * A more efficient way, binary search is used. It's not dp at all.
 * This way also cannot get the exact LIS.
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func lengthOfLIS2(nums []int) int {
	var LIS []int
	for i := 0; i < len(nums); i++ {
		lowerBoundIndex := lowerBound(LIS, nums[i])
		if lowerBoundIndex > len(LIS)-1 {
			LIS = append(LIS, nums[i])
		} else {
			LIS[lowerBoundIndex] = nums[i]
		}
	}
	return len(LIS)
}

func lowerBound(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
