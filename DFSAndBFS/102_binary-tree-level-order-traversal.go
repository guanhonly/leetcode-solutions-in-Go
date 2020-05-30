/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/binary-tree-level-order-traversal/
 */
package DFSAndBFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

/**
 * BFS is used here. Since there is no queue in Go's lib,
 * I used slice as a fake queue.
 * Time complexity: O(n), n is the num of nodes in a tree
 * Space complexity: O(n)
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q Queue
	queueAdd(&q, root)
	var res [][]int
	for len(q) > 0 {
		var oneLevel []int
		qSize := len(q)
		for i := 0; i < qSize; i++ {
			currNode := queuePop(&q)
			oneLevel = append(oneLevel, currNode.Val)
			if currNode.Left != nil {
				queueAdd(&q, currNode.Left)
			}
			if currNode.Right != nil {
				queueAdd(&q, currNode.Right)
			}
		}
		res = append(res, oneLevel)
	}
	return res
}

func queueAdd(q *Queue, val *TreeNode) {
	*q = append(*q, val)
}

func queuePop(q *Queue) *TreeNode {
	if len(*q) <= 0 {
		return nil
	}
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}
