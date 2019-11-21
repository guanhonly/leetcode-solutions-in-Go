/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/3sum/
 */
package HashTable

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo, hi, target := i+1, len(nums)-1, -nums[i]
			for lo < hi {
				if nums[lo]+nums[hi] < target {
					lo++
				} else if nums[lo]+nums[hi] > target {
					hi--
				} else {
					tmp := []int{nums[i], nums[lo], nums[hi]}
					res = append(res, tmp)
					for lo < hi && nums[lo+1] == nums[lo] {
						lo++
					}
					for lo < hi && nums[hi-1] == nums[hi] {
						hi--
					}
					lo++
					hi--
				}
			}
		}
	}
	return res
}
