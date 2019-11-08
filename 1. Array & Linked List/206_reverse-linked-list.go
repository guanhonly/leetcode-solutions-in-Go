/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/reverse-linked-list/
 */

package LinkedList

/** 
 * Iterative way. The moset efficient method.
 * Time complexity: O(n), n is the length of the linked list.
 * Space complexity: O(1)
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    var prevNode *ListNode
    currNode := head
    for currNode != nil {
        /**
         * a more concise way is:
         * currNode.Next, prevNode, currNode = prevNode, currNode, currNode.Next
        */
        nextNode := currNode.Next
        currNode.Next = prevNode
        prevNode = currNode
        currNode = nextNode
    }
    return prevNode
}

/**
 * Recursive way.
 * Time complexity: O(n)
 * Space complexity: O(1)
*/
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return p
}