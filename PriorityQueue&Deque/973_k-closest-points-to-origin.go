package PriorityQueueAndDeque

import "container/heap"

type Coor struct {
	X int
	Y int
}

type CoorHeap []*Coor

func (h CoorHeap) Len() int {
	return len(h)
}

func (h CoorHeap) Less(i, j int) bool {
	dis1 := h[i].X*h[i].X + h[i].Y*h[i].Y
	dis2 := h[j].X*h[j].X + h[j].Y*h[j].Y
	return dis1 < dis2
}

func (h CoorHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CoorHeap) Push(x interface{}) {
	*h = append(*h, x.(*Coor))
}

func (h *CoorHeap) Pop() interface{} {
	length := len(*h)
	x := (*h)[length-1]
	*h = (*h)[0 : length-1]
	return x
}

func kClosest(points [][]int, K int) [][]int {
	var res [][]int
	h := new(CoorHeap)
	heap.Init(h)
	for _, point := range points {
		heap.Push(h, &Coor{X: point[0], Y: point[1]})
	}
	for i := 0; i < K; i++ {
		coor := heap.Pop(h).(*Coor)
		res = append(res, []int{coor.X, coor.Y})
	}
	return res
}
