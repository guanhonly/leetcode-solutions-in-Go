/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/minimum-depth-of-binary-tree/
 */
package DFSAndBFS

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minDep := int(^uint(0) >> 1)
	if root.Left != nil {
		minDep = min(minDep, minDepth(root.Left))
	}
	if root.Right != nil {
		minDep = min(minDep, minDepth(root.Right))
	}
	return minDep + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
