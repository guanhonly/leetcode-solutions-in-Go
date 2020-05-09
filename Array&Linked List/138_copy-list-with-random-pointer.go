package LinkedList

/**
 * Definition for a Node.
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Copy the node and use the original node to link the copy
	// Before: o1->o2->o3->nil
	// After: o1->c1->o2->c2->o3->c3->nil
	curr := head
	for curr != nil {
		copyNode := &Node{
			Val: curr.Val,
		}
		next := curr.Next
		curr.Next = copyNode
		copyNode.Next = next
		curr = next
	}

	// Assign the Random of copy
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// Detach the copy from list
	dummy := new(Node)
	dummy.Next = head.Next
	newCurr := head.Next
	curr = head
	for newCurr.Next != nil {
		curr.Next = curr.Next.Next
		curr = curr.Next
		newCurr.Next = newCurr.Next.Next
		newCurr = newCurr.Next
	}
	curr.Next = nil
	return dummy.Next
}
