/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/number-of-1-bits/
 */
package BitwiseOperations

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
