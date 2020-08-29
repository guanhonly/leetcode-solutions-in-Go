package DynamicProgramming

func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m > n {
		return false
	}
	if m == 0 {
		return true
	}
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == m
}
