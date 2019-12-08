/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/triangle/description/
 */
package DynamicProgramming

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	min := triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min[j] = triangle[i][j] + intMin(min[j], min[j+1])
		}
	}
	return min[0]
}

func intMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
