package LinkedList

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle node
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the nodes after middle node
	mid := slow
	start := slow.Next
	for start.Next != nil {
		currNode := start.Next
		start.Next = currNode.Next
		currNode.Next = mid.Next
		mid.Next = currNode
	}

	// Reorder
	p1 := head
	p2 := mid.Next
	for p1 != mid {
		mid.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = mid.Next
	}
}
