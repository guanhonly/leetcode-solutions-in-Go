/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/implement-queue-using-stacks/
 */

package StackAndQueue

import "../utils"

/**
 * Another way may be better is to store the front value
 * of the queue, but it's not recommended that using variables
 * out of the function or struct, and the time complexity is
 * nearly the same as this way.
 * Time complexity: Push: O(1), Pop: Amortized O(1), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyQueue struct {
	Stack1, Stack2 *utils.Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		Stack1: utils.NewStack(),
		Stack2: utils.NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.Stack2.IsEmpty() {
		return this.Stack2.Pop().(int)
	}
	s1Size := this.Stack1.Size()
	for i := 0; i < s1Size-1; i++ {
		this.Stack2.Push(this.Stack1.Pop())
	}
	return this.Stack1.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.Stack2.IsEmpty() {
		return this.Stack2.Peek().(int)
	}
	if this.Stack1.IsEmpty() {
		return -1
	}
	s1Size := this.Stack1.Size()
	for i := 0; i < s1Size; i++ {
		this.Stack2.Push(this.Stack1.Pop())
	}
	return this.Stack2.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Stack1.IsEmpty() && this.Stack2.IsEmpty()
}

/**
 * This way may be impractical.
 * Time complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyQueue2 struct {
	Stack1, Stack2 *utils.Stack
}

/** Initialize your data structure here. */
func Constructor2() MyQueue2 {
	return MyQueue2{
		Stack1: utils.NewStack(),
		Stack2: utils.NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue2) Push(x int) {
	for !this.Stack1.IsEmpty() {
		this.Stack2.Push(this.Stack1.Pop())
	}
	this.Stack2.Push(x)
	for !this.Stack2.IsEmpty() {
		this.Stack1.Push(this.Stack2.Pop())
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue2) Pop() int {
	if this.Stack1.IsEmpty() {
		return -1
	}
	return this.Stack1.Pop().(int)
}

/** Get the front element. */
func (this *MyQueue2) Peek() int {
	return this.Stack1.Peek().(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue2) Empty() bool {
	return this.Stack1.IsEmpty()
}
