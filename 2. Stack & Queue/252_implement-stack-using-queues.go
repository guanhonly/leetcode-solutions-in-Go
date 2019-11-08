/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/implement-stack-using-queues/
 */

package StackAndQueue
/**
 * Since Queue is not supported in the Go's official lib,
 * I implement Queue here by using slice.
 */
type Queue struct {
	nums []int
}

func newQueue() *Queue {
	return &Queue{
		nums: []int{},
	}
}

func (q *Queue) push(n int) {
	q.nums = append(q.nums, n)
}

func (q *Queue) pop() int {
	if q.isEmpty() {
		return -1
	}
	res := q.nums[0]
	if len(q.nums) == 1 {
		q.nums = []int{}
	} else {
		q.nums = q.nums[1:]
	}
	return res
}
func (q *Queue) peek() int {
	if q.isEmpty() {
		return -1
	}
	return q.nums[0]
}

func (q *Queue) size() int {
	return len(q.nums)
}

func (q *Queue) isEmpty() bool {
	return q.size() == 0
}


/**
 * Two queues are used, there may be many other variants,
 * I just write my favorite one.
 * Time complexity: Push: O(1), Pop: O(n), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyStack struct {
	Queue1, Queue2 *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		Queue1: newQueue(),
		Queue2: newQueue(),
	}
}

var top int

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Queue1.push(x)
	top = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.Queue1.size() > 1 {
		top = this.Queue1.pop()
		this.Queue2.push(top)
	}
	res := this.Queue1.pop()
	this.Queue1, this.Queue2 = this.Queue2, this.Queue1
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue1.isEmpty()
}

/**
 * Only one queue is used, there may also be many other variants,
 * this is my favorite one, too.
 * Time complexity: Push: O(1), Pop: O(n), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyStack struct {
    Queue1, Queue2 *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{
        Queue1: newQueue(),
    }
}

var top int
/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Queue1.push(x)
    top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    originalSize := this.Queue1.size()
    for i:=0; i<originalSize-1; i++ {
        top = this.Queue1.pop()
        this.Queue1.push(top)
    }
    return this.Queue1.pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return this.Queue1.isEmpty()
}
