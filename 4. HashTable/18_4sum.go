/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/4sum/
 */
package HashTable

import (
	"sort"
)
/**
 * A customized way, not universal but efficient.
 * Time complexity: O(n^3)
 * Space complexity: O(1)
 */
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	maxNum := nums[n-1]
	if nums[0]*4 > target || maxNum*4 < target {
		return res
	}
	for i := 0; i < n-3; i++ {
		currNum := nums[i]
		if i > 0 && currNum == nums[i-1] {
			continue
		}
		if currNum+maxNum*3 < target {
			continue
		}
		if currNum*4 > target {
			break
		}
		if currNum*4 == target {
			if i+3 < n && nums[i+3] == currNum {
				tmp := []int{currNum, currNum, currNum, currNum}
				res = append(res, tmp)
			}
			break
		}
		threeSumForFourSum(i+1, n-1, target-currNum, currNum, nums, &res)
	}
	return res
}

func threeSumForFourSum(lo, hi, target, lastNum int, nums []int, res *[][]int) {
	if lo+1 >= hi {
		return
	}
	maxNum := nums[hi]
	if nums[lo]*3 > target || maxNum*3 < target {
		return
	}
	for i := lo; i < hi-1; i++ {
		currNum := nums[i]
        if i > lo && currNum == nums[i-1] {
            continue
        }
		if currNum+maxNum*2 < target {
			continue
		}
		if currNum*3 > target {
			return
		}
		if currNum*3 == target {
			if i+2 <= hi && nums[i+2] == currNum {
				tmp := []int{lastNum, currNum, currNum, currNum}
				*res = append(*res, tmp)
			}
            return
		}

		left, right, newTarget := i+1, hi, target-currNum
		for left < right {
			if nums[left]+nums[right] < newTarget {
				left++
			} else if nums[left]+nums[right] > newTarget {
				right--
			} else {
				tmp := []int{lastNum, currNum, nums[left], nums[right]}
				*res = append(*res, tmp)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
}

/**
 * A more universal way, to solve k-sum question.
 * Time complexity: O(n^3)
 * Space complexity: O(n)
 */
func fourSum2(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var res [][]int
	kSum(nums, target, 4, 0, &res, []int{})
	return res
}

func kSum(nums []int, target, k, start int, res *[][]int, path []int)  {
	n := len(nums)
	maxNum := nums[n-1]
	if n-start+1 < k || k < 2 || nums[start]*k > target || maxNum*k < target{
		return
	}
	if k == 2 {
		left, right := start, n-1
		for left < right {
			if nums[left] + nums[right] < target {
				left++
			} else if nums[left] + nums[right] > target {
				right--
			} else {
				*res = append(*res, append(path, nums[left], nums[right]))
				for left<right && nums[left+1] == nums[left] {
					left++
				}
				for left<right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			}
		}
	} else {
		for i:=start; i< n-k+1; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if nums[i] + maxNum*(k-1) < target {
				continue
			}
			if nums[i] * k > target {
				break
			}
			if nums[i] * k == target {
				if nums[i+k-1] == nums[i] {
					for j:=0; j<k; j++ {
						path = append(path, nums[i])
					}
					*res = append(*res, path)
				}
				break
			}
			kSum(nums, target-nums[i], k-1, i+1, res, append(path, nums[i]))
		}
	}
}

