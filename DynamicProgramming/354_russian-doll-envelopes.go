package DynamicProgramming

import "sort"

// Sort the width first, and take length as the LIS problem.
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})
	var dp []int
	for i := 0; i < len(envelopes); i++ {
		index := lowerBound(dp, envelopes[i][1])
		if index > len(dp)-1 {
			dp = append(dp, envelopes[i][1])
		} else {
			dp[index] = envelopes[i][1]
		}
	}
	return len(dp)
}
