package LinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	list = head
	return convert(listLength(head))
}

func listLength(head *ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}

func convert(length int) *TreeNode {
	if length == 0 {
		return nil
	}
	node := new(TreeNode)
	node.Left = convert(length / 2)
	node.Val = list.Val
	list = list.Next
	node.Right = convert(length - length/2 - 1)
	return node
}
