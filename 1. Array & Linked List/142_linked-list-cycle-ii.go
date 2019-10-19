/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/linked-list-cycle-ii
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * The similar way of linked-list-cycle's first solution.
 * Time complexity: O(n)
 * Space complexity: O(n)
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    visited := make(map[*ListNode]bool)
    for head != nil {
        if visited[head] {
            return head
        }
        visited[head] = true
        head = head.Next
    }
    return nil
}

/**
 * Implemented by Floyd algorithm.
 * This method can save memory space but I think it's
 * not the practical way to solve the problem since it
 * is harder to understand the principles than the
 * above method.
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            for head != fast {
                head = head.Next
                fast = fast.Next
            }
            return head
        }
    }
    return nil
}