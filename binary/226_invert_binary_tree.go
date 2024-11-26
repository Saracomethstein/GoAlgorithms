package binary

// link to the task on leetcode
// https://leetcode.com/problems/invert-binary-tree/description/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
