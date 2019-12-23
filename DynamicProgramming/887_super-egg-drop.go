/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/super-egg-drop/
 */

package DynamicProgramming

/**
 * DP. For dp[i][j], i is the number of eggs left,
 * j is the number of total floors. The transition
 * equation is:
 * dp[K][N] = min[1<=X<=N](max(dp[K-1][X-1], dp[K][N-X]) + 1)
 * Since the original dp is overtime, binary search
 * is used to optimize the dp.
 */
func superEggDrop(K int, N int) int {
    mem := make(map[[2]int]int)
	return dp(K, N, mem)
}
func dp(K, N int, mem map[[2]int]int) int {
    if N == 0 {
        return 0
    }
	if K == 1 {
		return N
	}

	if value, exist := mem[[2]int{K, N}]; exist {
		return value
	}
	left, right := 1, N
	res := int(^uint(0) >> 1)
	for left <= right {
		mid := (left + right) / 2
        eggBroken := dp(K-1, mid-1, mem)
		eggNotBroken := dp(K, N-mid, mem)

		if eggBroken > eggNotBroken {
			right = mid - 1
			res = min(res, eggBroken+1)
		} else {
		    left = mid + 1
		    res = min(res, eggNotBroken+1)
		}
	}
	mem[[2]int{K, N}] = res
	return res
}

/**
 * A more efficient way, for dp[k][m], k is
 * the number of eggs left, m is the maximum
 * of throwing eggs.
 */
func superEggDrop2(K int, N int) int {
    dp := make([][]int, K+1)
    for i:=0; i<len(dp); i++ {
        dp[i] = make([]int, N+1)
    }
    m := 0
    for dp[K][m] < N {
        m++
        for k:=1; k<=K; k++ {
            dp[k][m] = dp[k][m-1] + dp[k-1][m-1]+1
        }
    }
    return m
}

/**
 * dp compression of the method above.
 */
func superEggDrop3(K int, N int) int {
    dp := make([]int, K+1)
    m := 0
    for dp[K] < N {
        m++
        for i:=K; i>0; i-- {
            dp[i] = dp[i] + dp[i-1]+1
        }
    }
    return m
}