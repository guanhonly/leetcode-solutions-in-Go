package DynamicProgramming

func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        if i == 0 {
            dp[i][0] = grid[i][0]
        } else {
            dp[i][0] = dp[i-1][0]+grid[i][0]
        }
    }
    for i:=1; i<n; i++ {
        dp[0][i] = dp[0][i-1]+grid[0][i]
    }
    for i:=1; i<m; i++ {
        for j:=1; j<n; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[m-1][n-1]
}

func minPathSum2(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]int, n)
    dp[0] = grid[0][0]
    for i:=1; i<n; i++ {
        dp[i] = dp[i-1] + grid[0][i]
    }
    for i:=1; i<m; i++ {
        dp[0] += grid[i][0]
        for j:=1; j<n; j++ {
            dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
        }
    }
    return dp[n-1]
}
