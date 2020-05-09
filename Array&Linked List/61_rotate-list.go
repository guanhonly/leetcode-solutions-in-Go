package LinkedList

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	listLength := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		listLength++
	}

	// Construct a cycle
	tail.Next = head

	rotateIndex := listLength - k%listLength
	for i := 0; i < rotateIndex; i++ {
		tail = tail.Next
	}

	// Reach the rotate node and break the cycle
	head = tail.Next
	tail.Next = nil
	return head
}
