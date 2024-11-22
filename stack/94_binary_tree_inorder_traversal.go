package stack

// link to the task on leetcode
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := inorderTraversal(root.Left)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}
