package binary

import "fmt"

// link to the task on leetcode
// https://leetcode.com/problems/binary-tree-paths/description/

func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathStr(root, "")
}

func binaryTreePathStr(root *TreeNode, str string) []string {
	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		str += fmt.Sprint(root.Val)
		return []string{str}
	}
	str += fmt.Sprintf("%d->", root.Val)

	return append(
		binaryTreePathStr(root.Left, str),
		binaryTreePathStr(root.Right, str)...,
	)
}
