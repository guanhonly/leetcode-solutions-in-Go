package PriorityQueueAndDeque

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(NodeHeap)
	heap.Init(h)
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	dummyNode := new(ListNode)
	tail := dummyNode
	for h.Len() > 0 {
		head := heap.Pop(h).(*ListNode)
		tail.Next = head
		tail = tail.Next
		if head.Next != nil {
			heap.Push(h, head.Next)
		}
	}
	return dummyNode.Next
}
