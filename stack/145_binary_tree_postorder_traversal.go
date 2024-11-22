package stack

// link to the task on leetcode
// https://leetcode.com/problems/binary-tree-postorder-traversal/description/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := postorderTraversal(root.Left)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}
