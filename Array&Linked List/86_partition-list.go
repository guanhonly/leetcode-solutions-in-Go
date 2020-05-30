package LinkedList

func partition(head *ListNode, x int) *ListNode {
	lessHead, greaterHead := new(ListNode), new(ListNode)
	less, greater := lessHead, greaterHead
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}
	less.Next = greaterHead.Next
	greater.Next = nil
	return lessHead.Next
}
