/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/powx-n/
 */
package DivideAndConquer

/**
 * Recursive way.
 * Time complexity: O(logn)
 * Space complexity: O(logn)
 */
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return fastPow(x, int64(n))
}

func fastPow(x float64, n int64) float64 {
	if n == 0 {
		return 1
	}
	half := fastPow(x, n/2)
	if n&1 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

/**
 * Iterative way.
 * Time complexity: O(logn)
 * Space complexity: O(1)
 */
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	tmp := x
	for i := n; i > 0; i /= 2 {
		if i&1 == 1 {
			res *= tmp
		}
		tmp *= tmp
	}
	return res
}
