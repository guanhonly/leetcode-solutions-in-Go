package DynamicProgramming

func nthUglyNumber(n int) int {
	pos2, pos3, pos5 := 0, 0, 0
	k := make([]int, n)
	k[0] = 1
	for i := 1; i < n; i++ {
		k[i] = min(2*k[pos2], 3*k[pos3], 5*k[pos5])
		if k[i] == 2*k[pos2] {
			pos2++
		}
		if k[i] == 3*k[pos3] {
			pos3++
		}
		if k[i] == 5*k[pos5] {
			pos5++
		}
	}
	return k[n-1]
}
