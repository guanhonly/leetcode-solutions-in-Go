package BitwiseOperations

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m < n {
		m = m >> 1
		n = n >> 1
		shift++
	}
	return m << shift
}
