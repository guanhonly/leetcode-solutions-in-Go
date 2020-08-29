package String

func countBinarySubstrings(s string) int {
	var ans int
	prevChar, prevCount := s[0], 0
	for i := 0; i < len(s); {
		for ; i < len(s) && s[i] == prevChar; i++ {
			prevCount++
		}
		if i == len(s) {
			break
		}
		currChar := s[i]
		currCount := 0
		for ; i < len(s) && s[i] == currChar; i++ {
			currCount++
		}
		ans += min(prevCount, currCount)
		prevChar, prevCount = currChar, currCount
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
