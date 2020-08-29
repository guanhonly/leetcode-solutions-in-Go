package DynamicProgramming

func divisorGame(N int) bool {
	f := make([]bool, N+5)
	f[1], f[2] = false, true
	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !f[i-j] {
				f[i] = true
				break
			}
		}
	}
	return f[N]
}

// If N is even, Alice will win.
func divisorGame2(N int) bool {
	return N&1 == 0
}
