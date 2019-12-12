/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/lru-cache/
 */
package LRU

/**
 * Use double linked list and map.
 * Time complexity: Get:O(1), Put: O(1)
 * Space complexity: O(capacity)
 */
type DeLinkedListNode struct {
	key   int
	value int
	prev  *DeLinkedListNode
	next  *DeLinkedListNode
}
type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DeLinkedListNode
	head     *DeLinkedListNode
	tail     *DeLinkedListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.size = 0
	lru.capacity = capacity
	lru.cache = make(map[int]*DeLinkedListNode)
	lru.head = &DeLinkedListNode{}
	lru.tail = &DeLinkedListNode{}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.cache[key]; exist {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &DeLinkedListNode{
			key:   key,
			value: value,
		}
		this.addNode(newNode)
		this.cache[key] = newNode
		this.size++
		if this.size > this.capacity {
			lastNode := this.popTail()
			delete(this.cache, lastNode.key)
			this.size--
		}
	}
}

func (this *LRUCache) moveToHead(node *DeLinkedListNode) {
	this.detachNode(node)
	this.addNode(node)
}

func (this *LRUCache) detachNode(node *DeLinkedListNode) {
	prev := node.prev
	next := node.next
	prev.next, next.prev = next, prev
}

func (this *LRUCache) addNode(node *DeLinkedListNode) {
	node.prev, node.next = this.head, this.head.next
	this.head.next, this.head.next.prev = node, node
}

func (this *LRUCache) popTail() *DeLinkedListNode {
	lastNode := this.tail.prev
	this.detachNode(lastNode)
	return lastNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
