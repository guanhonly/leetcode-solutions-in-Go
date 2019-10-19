/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/linked-list-cycle
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * The most intuitive way, which is easy to think of,
 * but not sufficient enough.
 * Time complexity: O(n), n is the length of the linked list.
 * Space complexity: O(n)
 */
func hasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]bool)
    currNode := head
    for currNode != nil {
        if visited[currNode] {
            return true
        }
        visited[currNode] = true
        currNode = currNode.Next
    }
    return false
}

/**
 * This way may be anti-human but sufficient. Fast and slow
 * node can meet each other in a loop.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    fast, slow := head.Next, head
    for fast != slow {
        if fast == nil || fast.Next == nil {
            return false
        }
        fast = fast.Next.Next
        slow = slow.Next
    }
    return true
}