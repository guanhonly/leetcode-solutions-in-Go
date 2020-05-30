/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/valid-palindrome/
 */
package BinarySearch

// Since Go is not convenient as Java/C++/Python, some
// functions need to implemented by coder.
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		for left < len(s) && !isAlphaOrDigit(s[left]) {
			left++
		}
		for right >= 0 && !isAlphaOrDigit(s[right]) {
			right--
		}
		if left >= right {
			break
		}
		leftChar := toLower(s[left])
		rightChar := toLower(s[right])
		if leftChar != rightChar {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphaOrDigit(a byte) bool {
	if a >= 48 && a <= 57 {
		return true
	}
	if a >= 65 && a <= 90 {
		return true
	}
	if a >= 97 && a <= 122 {
		return true
	}
	return false
}

func toLower(a byte) byte {
	if a >= 65 && a <= 90 {
		a += 32
	}
	return a
}
