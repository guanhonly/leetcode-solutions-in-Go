package BinaryTree

var f, g map[*TreeNode]int

func rob(root *TreeNode) int {
	f, g = make(map[*TreeNode]int), make(map[*TreeNode]int)
	dfs(root)
	return max(f[root], g[root])
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	f[root] = root.Val + g[root.Left] + g[root.Right]
	g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
}
