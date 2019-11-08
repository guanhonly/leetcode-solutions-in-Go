/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/swap-nodes-in-pairs
 */

package LinkedList
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
    Val int
    Next *ListNode
}


/**
 * The recursive way, the code is clear but not easy
 * to think it clear since recursion is anti-human.
 * Don't get caught in the details of recursion since
 * it's the machine's business, just think that if you
 * want to swap current pair of nodes, the head.Next
 * must point to the node that swapped, then let nextNode
 * point to head. The abort condition of recursion is
 * the node is nil, which means the final node is reached.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    nextNode := head.Next
    head.Next = swapPairs(nextNode.Next)
    nextNode.Next = head
    return nextNode
}

/**
 * The non-recursive way. The code is more complex but
 * easier to understand than recursive way.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func swapPairs2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummyNode := &ListNode{Val:0}
    dummyNode.Next = head
    currNode := dummyNode
    for currNode.Next != nil && currNode.Next.Next != nil {
        start := currNode.Next
        end := currNode.Next.Next
        start.Next = end.Next
        end.Next = start
        currNode.Next = end
        currNode = start
    }
    return dummyNode.Next
}