/**
 * Difficulty: Easy
 * Question link:  https://leetcode.com/problems/lemonade-change/
 */

package GreedyAlgorithm

func lemonadeChange(bills []int) bool {
	fiveCount := 0
	tenCount := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			fiveCount++
		} else if bills[i] == 10 {
			if fiveCount <= 0 {
				return false
			}
			fiveCount--
			tenCount++
		} else if bills[i] == 20 {
			if tenCount > 0 && fiveCount > 0 {
				tenCount--
				fiveCount--
			} else if fiveCount >= 3 {
				fiveCount -= 3
			} else {
				return false
			}
		}
	}
	return true
}
