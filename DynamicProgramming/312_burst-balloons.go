package DynamicProgramming

// DP records the coins in range (i, j).
var vals []int
var dpArray [][]int
func maxCoins(nums []int) int {
    n := len(nums)
    vals = make([]int, n+2)
    vals[0] = 1
    vals[n+1] = 1
    for i:=1; i<=n; i++ {
        vals[i] = nums[i-1]
    }
    dpArray = make([][]int, n+2)
    for i:=0; i<n+2; i++ {
        dpArray[i] = make([]int, n+2)
        for j:=0; j<n+2; j++ {
            dpArray[i][j] = -1
        }
    }
    return solve(0, n+1)
}

func solve(left, right int) int {
    if left >= right -1 {
        return 0
    }
    if dpArray[left][right] != -1 {
        return dpArray[left][right]
    }
    for i:=left+1; i<right; i++ {
        coins := vals[left] * vals[i] * vals[right]
        coins += solve(left, i) + solve(i, right)
        dpArray[left][right] = max(dpArray[left][right], coins)
    }
    return dpArray[left][right]
}


// Pure DP.
func maxCoins2(nums []int) int {
    n := len(nums)
    vals := make([]int,0, n+2)
    vals = append([]int{1}, nums...)
    vals = append(vals, 1)
    dp := make([][]int, n+2)
    for i:=0; i<n+2; i++ {
        dp[i] = make([]int, n+2)
    }
    for i:=n-1; i>=0; i-- {
        for j:=i+2; j<n+2; j++ {
            for k:=i+1; k<j; k++ {
                coins := vals[i] * vals[k] * vals[j]
                coins += dp[i][k] + dp[k][j]
                dp[i][j] = max(dp[i][j], coins)
            }
        }
    }
    return dp[0][n+1]
}