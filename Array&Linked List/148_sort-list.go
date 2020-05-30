package LinkedList

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	midNode := findMid(head)
	rightNode := sortList(midNode.Next)
	midNode.Next = nil
	leftNode := sortList(head)
	return merge(leftNode, rightNode)
}

func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}
