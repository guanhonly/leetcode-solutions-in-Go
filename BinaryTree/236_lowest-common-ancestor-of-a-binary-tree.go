/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
 */
package BinaryTree

/**
 * Recursion is used, to search the left subtree and right subtree recursively,
 * until find the node that is p or q or their common parent.
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor3(root.Left, p, q)
	right := lowestCommonAncestor3(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	// If p and q are in the left and right subtrees, return the root
	return root
}
