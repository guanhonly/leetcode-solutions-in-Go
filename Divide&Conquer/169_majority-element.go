/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/majority-element/
 */
package DivideAndConquer

import "sort"

/**
 * Divide and conquer.
 * Time complexity: O(nlogn), n is the length of nums
 * Space complexity: O(logn)
 */
func majorityElement(nums []int) int {
	return majorityElementInRange(nums, 0, len(nums)-1)
}

func majorityElementInRange(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}
	mid := (hi-lo)/2 + lo
	leftMajority := majorityElementInRange(nums, lo, mid)
	rightMajority := majorityElementInRange(nums, mid+1, hi)
	if leftMajority == rightMajority {
		return leftMajority
	}
	leftCount := countInRange(nums, leftMajority, lo, hi)
	rightCount := countInRange(nums, rightMajority, lo, hi)
	if leftCount > rightCount {
		return leftMajority
	} else {
		return rightMajority
	}
}

func countInRange(nums []int, target, lo, hi int) int {
	count := 0
	for i := lo; i <= hi; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

/**
 * Sort the array, then the num in the middle is the majority num.
 * Time complexity: O(nlogn)
 * Space complexity: O(1)
 */
func majorityElement2(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)/2]
}

/**
 * Boyer-Moore Voting Algorithm
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func majorityElement3(nums []int) int {
	count := 0
	var majority int
	for _, i := range nums {
		if count == 0 {
			majority = i
		}
		if i == majority {
			count++
		} else {
			count--
		}
	}
	return majority
}
