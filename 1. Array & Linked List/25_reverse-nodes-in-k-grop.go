/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/reverse-nodes-in-k-group/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * Non-recursive way.
 * Take k nodes as a trunk, detach a trunk from the list then reverse
 *  the nodes inside a trunk. Finally reconnect the trunk to the list,
 * just take care of the prevNode and nextNode of the trunk.
 * Time complexity: O(n*k), n is the length of the list, best O(n), worst O(n^2)
 * Space complexity: O(1)
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummyNode := &ListNode{Val:0}
    dummyNode.Next = head
    prevNode, endNode := dummyNode, dummyNode
    for endNode.Next != nil {
        for i := 0; i<k && endNode != nil; i++ {
            endNode = endNode.Next
        }
        if endNode == nil {
            break
        }
        startNode := prevNode.Next
        nextNode := endNode.Next
        endNode.Next = nil
        prevNode.Next = reverseList(startNode)
        startNode.Next = nextNode
        prevNode = startNode
        endNode = startNode
    }
    return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var prevNode *ListNode = nil
    currNode := head
    for currNode != nil {
        currNode.Next, prevNode, currNode = prevNode, currNode, currNode.Next
    }
    return prevNode
}

/**
 * Recursive way, similar to the recursive way in swap-nodes-in-pairs.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    prevNode, endNode := head, head
    for i:= 0; i<k-1 && endNode != nil; i++ {
        endNode = endNode.Next
    }
    if endNode == nil {
        return head
    }
    nextNode := endNode.Next
    endNode.Next = nil
    reversedHead := reverseList(prevNode)
    if nextNode != nil {
        prevNode.Next = reverseKGroup(nextNode, k)
    }
    return reversedHead
}