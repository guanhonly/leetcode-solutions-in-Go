package DynamicProgramming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 {
        return 0
    }
    dp := make([][]int, m)
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
        for j:=0; j<n; j++ {
            if obstacleGrid[i][j] == 0 {
                dp[i][j] = 1
            } else {
                dp[i][j] = 0
            }
        }
    }
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            if i >= 1 && j >= 1 {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            } else if i >= 1 {
                dp[i][j] = dp[i-1][j]
            } else if j >= 1 {
                dp[i][j] = dp[i][j-1]
            }
        }
    }
    return dp[m-1][n-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obstacleGrid[0])
    dp := make([]int, n)
    if obstacleGrid[0][0] == 0 {
        dp[0] = 1
    } else {
        dp[0] = 0
    }
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
                continue
            }
            if j>=1 && obstacleGrid[i][j-1] == 0 {
                dp[j] += dp[j-1]
            }
        }
    }
    return dp[n-1]
}