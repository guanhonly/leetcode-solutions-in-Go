package PriorityQueueAndDeque

import "container/heap"

type TupleHeap []Tuple

func (h TupleHeap) Len() int           { return len(h) }
func (h TupleHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h TupleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TupleHeap) Push(x interface{}) {
	*h = append(*h, x.(Tuple))
}

func (h *TupleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := new(TupleHeap)
	heap.Init(h)
	for i := 0; i < n; i++ {
		heap.Push(h, Tuple{X: 0, Y: i, Val: matrix[0][i]})
	}
	for i := 0; i < k-1; i++ {
		val := heap.Pop(h).(Tuple)
		if val.X == n-1 {
			continue
		}
		heap.Push(h, Tuple{X: val.X + 1, Y: val.Y, Val: matrix[val.X+1][val.Y]})
	}
	return heap.Pop(h).(Tuple).Val
}

type Tuple struct {
	X   int
	Y   int
	Val int
}
