package BinaryTree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	if root.Left == nil && root.Right == nil {
		paths = append(paths, strconv.Itoa(root.Val))
		return paths
	}

	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)
	for _, path := range leftPaths {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+path)
	}
	for _, path := range rightPaths {
		paths = append(paths, strconv.Itoa(root.Val)+"->"+path)
	}
	return paths
}

// Another solution, use a void func. But I don't
// like this kind of code style, since a slice pointer
// *[]string is used here.
func binaryTreePaths2(root *TreeNode) []string {
	var res []string
	if root != nil {
		helper2(root, "", &res)
	}
	return res
}

func helper2(root *TreeNode, path string, paths *[]string) {
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path+strconv.Itoa(root.Val))
	}
	if root.Left != nil {
		helper2(root.Left, path+strconv.Itoa(root.Val)+"->", paths)
	}
	if root.Right != nil {
		helper2(root.Right, path+strconv.Itoa(root.Val)+"->", paths)
	}
}
