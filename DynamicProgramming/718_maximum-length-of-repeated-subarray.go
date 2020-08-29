package DynamicProgramming

func findLength(A []int, B []int) int {
	lenOfA := len(A)
	lenOfB := len(B)
	dp := make([][]int, lenOfA+1)
	for i := range dp {
		dp[i] = make([]int, lenOfB+1)
	}
	ans := 0
	for i := lenOfA - 1; i >= 0; i-- {
		for j := lenOfB - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

func findLength2(A []int, B []int) int {
	lenOfA := len(A)
	lenOfB := len(B)
	dp := make([]int, lenOfB)
	ans := 0
	for i := lenOfA - 1; i >= 0; i-- {
		prev := 0
		for j := lenOfB - 1; j >= 0; j-- {
			ori := dp[j]
			if A[i] == B[j] {
				dp[j] = prev + 1
			} else {
				dp[j] = 0
			}
			prev = ori
			ans = max(ans, dp[j])
		}
	}
	return ans
}
