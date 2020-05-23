package TwoPointers

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return closestSum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
