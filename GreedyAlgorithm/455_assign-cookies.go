/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/assign-cookies/
 */
package GreedyAlgorithm

import "sort"

/**
 * First meet the need of child with big greed factor.
 * Time complexity: O(nlogn + n), n is the longer of length
 * of g and length of s.
 * Space complexity: O(1)
 */
func findContentChildren(g []int, s []int) int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	count := 0
	for i := 0; i < len(g) && count < len(s); i++ {
		if s[count] >= g[i] {
			count++
		}
	}
	return count
}

/**
 * First meet the need of child with small greed factor.
 * Time complexity: O(nlogn + n)
 * Space complexity: O(1)
 */
func findContentChildren2(g []int, s []int) int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	res := 0
	gIndex := 0
	sIndex := 0
	for gIndex < len(g) && sIndex < len(s) {
		if s[sIndex] >= g[gIndex] {
			res++
			gIndex++
		}
		sIndex++
	}
	return res
}
