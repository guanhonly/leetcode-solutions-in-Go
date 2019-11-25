/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
 */
package DFSAndBFS

/**
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1+max(maxDepth(root.Left), maxDepth(root.Right))
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
