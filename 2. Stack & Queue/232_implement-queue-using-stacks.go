/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/implement-queue-using-stacks/
 */

/**
 * Since Stack is not supported in the Go's official lib,
 * I implement Stack here by using slice.
 */
type Stack struct {
	nums []int
}

func newStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

func (s *Stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}
func (s *Stack) peek() int {
	if s.isEmpty() {
		return -1
	}
	return s.nums[len(s.nums)-1]
}

func (s *Stack) size() int {
	return len(s.nums)
}

func (s *Stack) isEmpty() bool {
	return s.size() == 0
}


/**
 * Another way may be better is to store the front value
 * of the queue, but it's not recommended that using variables
 * out of the function or struct, and the time complexity is
 * nearly the same as this way.
 * Time complexity: Push: O(1), Pop: Amortized O(1), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyQueue struct {
	Stack1, Stack2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack1: newStack(),
		Stack2: newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack1.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.Stack2.isEmpty() {
		return this.Stack2.pop()
	}
	s1Size := this.Stack1.size()
	for i := 0; i < s1Size-1; i++ {
		this.Stack2.push(this.Stack1.pop())
	}
	return this.Stack1.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.Stack2.isEmpty() {
		return this.Stack2.peek()
	}
	if this.Stack1.isEmpty() {
		return -1
	}

	return this.Stack1.nums[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Stack1.isEmpty() && this.Stack2.isEmpty()
}


/**
 * This way may be impractical.
 * Time complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyQueue struct {
	Stack1, Stack2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack1: newStack(),
		Stack2: newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for !this.Stack1.isEmpty() {
		this.Stack2.push(this.Stack1.pop())
	}
	this.Stack2.push(x)
	for !this.Stack2.isEmpty() {
		this.Stack1.push(this.Stack2.pop())
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Stack1.isEmpty() {
		return -1
	}
	return this.Stack1.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Stack1.peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Stack1.isEmpty()
}
