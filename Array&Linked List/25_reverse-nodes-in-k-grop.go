/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/reverse-nodes-in-k-group/
 */

package LinkedList

/**
 * Non-recursive way.
 * Take k nodes as a trunk, detach a trunk from the list then reverse
 *  the nodes inside a trunk. Finally reconnect the trunk to the list,
 * just take care of the prevNode and nextNode of the trunk.
 * Time complexity: O(n*k), n is the length of the list, best O(n), worst O(n^2)
 * Space complexity: O(1)
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	prev := dummy
	for prev != nil {
		prev = reverseNextKNodes(prev, k)
	}
	return dummy.Next
}

func reverseNextKNodes(head *ListNode, k int) *ListNode {
	currNode := head
	for i := 0; i < k; i++ {
		currNode = currNode.Next
		if currNode == nil {
			return currNode
		}
	}
	firstNode := head.Next
	lastNode := currNode

	nextNode := currNode.Next

	// reverse k nodes
	prevNode := head
	currNode = head.Next
	for currNode != nextNode {
		currNode, currNode.Next, prevNode = currNode.Next, prevNode, currNode
	}

	// connect to the original node list
	head.Next = lastNode
	firstNode.Next = nextNode
	return firstNode
}

/**
 * Recursive way, similar to the recursive way in swap-nodes-in-pairs.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func reverseKGroup2(head *ListNode, k int) *ListNode {
	prevNode, endNode := head, head
	for i := 0; i < k-1 && endNode != nil; i++ {
		endNode = endNode.Next
	}
	if endNode == nil {
		return head
	}
	nextNode := endNode.Next
	endNode.Next = nil
	reversedHead := reverseList(prevNode)
	if nextNode != nil {
		prevNode.Next = reverseKGroup2(nextNode, k)
	}
	return reversedHead
}
