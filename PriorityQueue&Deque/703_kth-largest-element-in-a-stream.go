/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/kth-largest-element-in-astream/
 */

package PriorityQueueAndDeque

type IntHeap []int
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Peek() interface{} {
	return (*h)[0]
}


type KthLargest struct {
    k int
    myHeap IntHeap
}


func Constructor(k int, nums []int) KthLargest {
	var myHeap IntHeap
	heap.Init(&myHeap)
	h := KthLargest{
		k:      k,
		myHeap: myHeap,
	}
	for _, i := range nums {
		h.Add(i)
	}
	return h
}


func (this *KthLargest) Add(val int) int {
	if this.myHeap.Len() < this.k {
		heap.Push(&this.myHeap, val)
	} else if val > this.myHeap.Peek().(int) {
		heap.Pop(&this.myHeap)
		heap.Push(&this.myHeap, val)
	}
	return this.myHeap.Peek().(int)
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */