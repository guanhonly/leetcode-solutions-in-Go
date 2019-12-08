/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/sqrtx/
 */

package BinarySearch

/**
 * The binary search way, a trick is that the
 * initial value of right is x/2+1, since the
 * sqrt of a positive number cannot be greater
 * than the 1/2 of itself.
 * Time complexity: O(logx)
 * Space complexity: O(1)
 */
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x/2+1
	var res int
	for left <= right {
		mid := (left + right) / 2
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
			res = mid
		}
	}
	return res
}

/**
 * Newton method, which is a convenient way to find root
 * of an equation.
 * Time complexity: O(logx)
 * Space complexity: O(1)
 */
func mySqrt2(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
