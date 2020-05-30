/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/sum-of-two-integers/
 */
package BitwiseOperations

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
