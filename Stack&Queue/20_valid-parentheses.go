/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/valid-parentheses/
 */
package StackAndQueue

import "../utils"

/**
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func isValid(s string) bool {
	if (len(s) & 1) > 0 {
		return false
	}
	stack := utils.NewStack()
	matching := make(map[int32]int32)
	matching[')'] = '('
	matching[']'] = '['
	matching['}'] = '{'
	for _, c := range s {
		if _, ok := matching[c]; !ok {
			stack.Push(int32(c))
		} else {
			if stack.IsEmpty() {
				return false
			}
			if !(matching[c] == stack.Pop()) {
				return false
			}
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}
