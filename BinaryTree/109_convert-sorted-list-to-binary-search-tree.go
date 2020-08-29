package BinaryTree

type ListNode struct {
	Val  int
	Next *ListNode
}

var globalHead *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	globalHead = head
	length := getLength(head)
	return buildTree(0, length-1)
}

func getLength(head *ListNode) int {
	ret := 0
	for ; head != nil; head = head.Next {
		ret++
	}
	return ret
}

func buildTree(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2
	root := &TreeNode{}
	root.Left = buildTree(left, mid-1)
	root.Val = globalHead.Val
	globalHead = globalHead.Next
	root.Right = buildTree(mid+1, right)
	return root
}
