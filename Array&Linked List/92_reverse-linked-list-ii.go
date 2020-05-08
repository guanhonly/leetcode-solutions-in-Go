package LinkedList

// A straightforward way, code is long but easy to understand.
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil {
        return head
    }
    dummy := new(ListNode)
    dummy.Next = head
    var prev *ListNode
    for i:=0; i<m-1; i++ {
        prev = head
        head = head.Next
    }
    first := head
    var prevNode *ListNode

    // Reverse
    for i:=0; i<n-m; i++ {
        tmp := head.Next
        head.Next = prevNode
        prevNode = head
        head = tmp
    }

    // Now the head is the last node that needs reverse
    nextNode := head.Next
    head.Next = prevNode
    if prev != nil {
        prev.Next = head
    }
    first.Next = nextNode

    // Edge case, if m == 1, dummy.Next won't return the real value.
    if prev == nil {
        return head
    }
    return dummy.Next
}

// This concise way is great but a little hard to understand.
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
    if head == nil {
        return head
    }
    dummy := new(ListNode)
    dummy.Next = head
    prev := dummy
    for i:=0; i<m-1; i++ {
        prev = prev.Next
    }
    start := prev.Next
    then := start.Next

    for i:=0; i<n-m; i++ {
        start.Next = then.Next
        then.Next = prev.Next
        prev.Next = then
        then = start.Next
    }
    return dummy.Next
}