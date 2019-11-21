/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/implement-stack-using-queues/
 */
package StackAndQueue

import "../utils"

/**
 * Two queues are used, there may be many other variants,
 * I just write my favorite one.
 * Time complexity: Push: O(1), Pop: O(n), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyStack struct {
	Queue1, Queue2 *utils.Queue
}

/** Initialize your data structure here. */
func Constructor3() MyStack {
	return MyStack{
		Queue1: utils.NewQueue(),
		Queue2: utils.NewQueue(),
	}
}

var top int

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Queue1.Push(x)
	top = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.Queue1.Size() > 1 {
		top = this.Queue1.Pop().(int)
		this.Queue2.Push(top)
	}
	res := this.Queue1.Pop().(int)
	this.Queue1, this.Queue2 = this.Queue2, this.Queue1
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue1.IsEmpty()
}

/**
 * Only one queue is used, there may also be many other variants,
 * this is my favorite one, too.
 * Time complexity: Push: O(1), Pop: O(n), Peek: O(1), Empty: O(1)
 * Space complexity: Push: O(n), Pop: O(1), Peek: O(1), Empty: O(1)
 */
type MyStack2 struct {
	Queue1, Queue2 *utils.Queue
}

/** Initialize your data structure here. */
func Constructor4() MyStack2 {
	return MyStack2{
		Queue1: utils.NewQueue(),
	}
}

var top2 int

/** Push element x onto stack. */
func (this *MyStack2) Push(x int) {
	this.Queue1.Push(x)
	top = x
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack2) Pop() int {
	originalSize := this.Queue1.Size()
	for i := 0; i < originalSize-1; i++ {
		top = this.Queue1.Pop().(int)
		this.Queue1.Push(top)
	}
	return this.Queue1.Pop().(int)
}

/** Get the top element. */
func (this *MyStack2) Top() int {
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack2) Empty() bool {
	return this.Queue1.IsEmpty()
}
