package HashTable

func palindromePairs(words []string) [][]int {
	wordsRev := []string{}
	indices := map[string]int{}

	n := len(words)
	for _, word := range words {
		wordsRev = append(wordsRev, reverse(word))
	}
	for i := 0; i < n; i++ {
		indices[wordsRev[i]] = i
	}

	ret := [][]int{}
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		if m == 0 {
			continue
		}
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1, indices)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1, indices)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}
