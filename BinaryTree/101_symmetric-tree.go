/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/symmetric-tree/
 */

package BinaryTree

func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}

func isMirror(r1, r2 *TreeNode) bool {
    if r1 == nil && r2 == nil {
        return true
    }
    if r1 == nil || r2 == nil {
        return false
    }
    return (r1.Val == r2.Val) &&
        isMirror(r1.Left, r2.Right) &&
        isMirror(r1.Right, r2.Left)
}

func isSymmetric2(root *TreeNode) bool {
    var queue []*TreeNode
    queue = append(queue, root)
    queue = append(queue, root)
    for len(queue) > 0 {
        r1 := queue[0]
        r2 := queue[1]
        queue = queue[2:]
        if r1 == nil && r2 == nil {
            continue
        }
        if r1 == nil || r2 == nil {
            return false
        }
        if r1.Val != r2.Val {
            return false
        }
        queue = append(queue, r1.Left)
        queue = append(queue, r2.Right)
        queue = append(queue, r1.Right)
        queue = append(queue, r2.Left)
    }
    return true
}