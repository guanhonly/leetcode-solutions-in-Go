package DFSAndBFS

func partition(s string) [][]string {
	var res [][]string
	if len(s) == 0 {
		return res
	}

	var palindromes []string
	helper(0, s, &palindromes, &res)
	return res
}

func helper(startIndex int, s string, palindromes *[]string, res *[][]string) {
	if startIndex == len(s) {
		tmp := make([]string, len(*palindromes))
		copy(tmp, *palindromes)
		*res = append(*res, tmp)
	}
	for i := startIndex; i < len(s); i++ {
		subString := s[startIndex : i+1]
		if !isPalindrome(subString) {
			continue
		}
		*palindromes = append(*palindromes, subString)
		helper(i+1, s, palindromes, res)
		*palindromes = (*palindromes)[:len(*palindromes)-1]
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	left, right := 0, n-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Another trick, optimize the palindrome validation.
var palindromeValidation [][]bool
var results [][]string

func partition2(s string) [][]string {
	results = [][]string{}
	if len(s) == 0 {
		return results
	}

	palindromeValidation = getPalindromeValidation(s)
	var cutIndex []int
	helper2(0, s, &cutIndex)
	return results
}

func helper2(startIndex int, s string, cutIndex *[]int) {
	if startIndex == len(s) {
		addResult(s, *cutIndex)
	}
	for i := startIndex; i < len(s); i++ {
		if !palindromeValidation[startIndex][i] {
			continue
		}
		*cutIndex = append(*cutIndex, i)
		helper2(i+1, s, cutIndex)
		*cutIndex = (*cutIndex)[:len(*cutIndex)-1]
	}
}

func getPalindromeValidation(s string) [][]bool {
	n := len(s)
	res := make([][]bool, n)
	for i := range res {
		res[i] = make([]bool, n)
		res[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		res[i][i+1] = s[i] == s[i+1]
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			res[i][j] = res[i+1][j-1] && s[i] == s[j]
		}
	}
	return res
}

func addResult(s string, cutIndex []int) {
	var res []string
	startIndex := 0
	for i := 0; i < len(cutIndex); i++ {
		res = append(res, s[startIndex:cutIndex[i]+1])
		startIndex = cutIndex[i] + 1
	}
	results = append(results, res)
}
