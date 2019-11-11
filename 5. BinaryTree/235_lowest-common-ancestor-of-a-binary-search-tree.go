/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
 */
package BinaryTree

/**
 * Recursive way.
 * Time complexity: O(n), n is the number of nodes in the binary tree
 * Space complexity: O(n), since recursion is used, but the space usage
 * won't increase too much since only pointers are used in recursion.
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

/**
 * Non-recursive way.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	node := root
	for node != nil {
		if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		} else if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else {
			break
		}
	}
	return node
}
