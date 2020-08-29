package DynamicProgramming

func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var calcPoints func(l, r, k int) int
	calcPoints = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		for r > l && boxes[r] == boxes[r-1] {
			r--
			k++
		}
		dp[l][r][k] = calcPoints(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = max(dp[l][r][k], calcPoints(l, i, k+1)+calcPoints(i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return calcPoints(0, len(boxes)-1, 0)
}
