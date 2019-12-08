/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/climbing-stairs/
 */
package DynamicProgramming

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}
	before := 1
	allWays := 2
	for i := 2; i < n; i++ {
		allWays, before = allWays+before, allWays
	}
	return allWays
}
