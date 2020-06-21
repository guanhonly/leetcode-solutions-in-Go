package BinaryTree

const INT_MIN = ^(int(^uint(0) >> 1))

func maxPathSum(root *TreeNode) int {
	res := INT_MIN
	maxGain(root, &res)
	return res
}

func maxGain(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	leftGain := max(0, maxGain(root.Left, res))
	rightGain := max(0, maxGain(root.Right, res))

	currMax := root.Val + leftGain + rightGain
	*res = max(*res, currMax)
	return root.Val + max(leftGain, rightGain)
}
