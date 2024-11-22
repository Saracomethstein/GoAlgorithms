package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// link to the task on leetcode
// https://leetcode.com/problems/binary-tree-preorder-traversal/description/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{root.Val}
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}
