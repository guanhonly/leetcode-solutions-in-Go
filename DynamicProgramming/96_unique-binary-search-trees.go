package DynamicProgramming

// 如果i是根，则前i-1个数能构成的BST数量和后n-i个数能构成的BST数量即为以i为根的BST数量
// 设以i为根，长度为n的BST数量为F(i, n), 长度为n的BST数量为G(n), 则
// F(i, n) = G(i-1)*G(n-i)
// G(n) = sum(F(i, n)), i=1,..., n
// 即G(n) = sum(G(i-1)*G(n-i)), i=1,...,n
func numTrees(n int) int {
    G := make([]int, n+1)
    G[0] = 1
    G[1] = 1
    for i:=2; i<=n; i++ {
        for j:=1; j<=i; j++ {
            G[i] += G[j-1]*G[i-j]
        }
    }
    return G[n]
}

// Catalan
func numTrees2(n int) int {
    C := 1
    for i:=0; i<n; i++ {
        C = C * 2 * (2 * i + 1) / (i + 2)
    }
    return C
}