package DynamicProgramming

// Common dp.
func isMatch(s string, p string) bool {
    m := len(s)
    n := len(p)
    dp := make([][]bool, m+1)
    for i:=0; i<=m; i++ {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for i:=0; i<n; i++ {
        if p[i] == '*' {
            dp[0][i+1] = true
        } else {
            break
        }
    }
    for i:=1; i<=m; i++ {
        for j:=1; j<=n; j++ {
            if p[j-1] == '?' || s[i-1] == p[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i][j-1] || dp[i-1][j]
            }
        }
    }
    return dp[m][n]
}

// Optimized dp.
func isMatch2(s string, p string) bool {
    previ, prevj := -1, -1
    m := len(s)
    n := len(p)
    i, j := 0, 0
    for i<m {
        if j<n && (s[i] == p[j] || p[j] == '?') {
            i++
            j++
        } else if j<n && p[j] == '*' {
            previ = i
            prevj = j
            j++
        } else if previ >= 0 {
            i = previ + 1
            previ++
            j = prevj+1
        } else {
            return false
        }
    }
    for j<n && p[j] == '*' {
        j++
    }
    return j == n
}