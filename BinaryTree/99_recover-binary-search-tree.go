package BinaryTree

func recoverTree(root *TreeNode) {
	var nums []int
	inOrder(root, &nums)
	x, y := findTwoAbnormalVals(nums)
	recover(root, 2, x, y)
}

func inOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inOrder(root.Right, nums)
}

func findTwoAbnormalVals(nums []int) (int, int) {
	n := len(nums)
	x, y := -1, -1
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root != nil {
		if root.Val == x {
			root.Val = y
			count--
		} else if root.Val == y {
			root.Val = x
			count--
		}
		if count == 0 {
			return
		}
		recover(root.Right, count, x, y)
		recover(root.Left, count, x, y)
	}
}
