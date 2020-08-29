package BinaryTree

import "math"

func isBalanced(root *TreeNode) bool {
	_, balanced := helper3(root)
	return balanced
}

func helper3(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalanced := helper3(root.Left)
	rightDepth, rightBalanced := helper3(root.Right)

	if !leftBalanced || !rightBalanced {
		return -1, false
	}
	if rightDepth-leftDepth > 1 || leftDepth-rightDepth > 1 {
		return -1, false
	}
	return 1 + max(leftDepth, rightDepth), true
}

// In helper4, return -1 means that it's not a
// balanced tree, otherwise the return value is
// the depth of the subtree. This is more concise
// but the code style is not very well.
func isBalanced2(root *TreeNode) bool {
	return helper4(root) != -1
}

func helper4(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	leftDepth := helper4(root.Left)
	rightDepth := helper4(root.Right)
	if leftDepth == -1 || rightDepth == -1 || math.Abs(leftDepth-rightDepth) > 1 {
		return -1
	}
	return math.Max(leftDepth, rightDepth) + 1
}
