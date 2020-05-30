package TwoPointers

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right {
		if !isLetterOrDigit(s[left]) {
			left++
		} else if !isLetterOrDigit(s[right]) {
			right--
		} else {
			if toLower(s[left]) != toLower(s[right]) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isLetterOrDigit(char byte) bool {
	if char >= 48 && char <= 57 {
		return true
	}

	if char >= 65 && char <= 90 {
		return true
	}

	if char >= 97 && char <= 122 {
		return true
	}
	return false
}

func toLower(char byte) byte {
	if char >= 65 && char <= 90 {
		return char + 32
	}
	return char
}
