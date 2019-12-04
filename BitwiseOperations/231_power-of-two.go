/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/power-of-two/
 */
package BitwiseOperations

/**
 * The bruce force solution.
 * Time complexity: O(logn)
 * Space complexity: O(1)
 */
func isPowerOfTwo(n int) bool {
    init := 1
    for n >= init {
        if n == init {
            return true
        }
        init = init<<1
    }
    return false
}

/**
 * If a number n is a power of two, n&(n-1) must equal
 * to 0.
 * Time complexity: O(1)
 * Space complexity: O(1)
 */
func isPowerOfTwo2(n int) bool {
    if n <= 0 {
        return false
    }
    if n&(n-1) == 0 {
        return true
    }
    return false
}