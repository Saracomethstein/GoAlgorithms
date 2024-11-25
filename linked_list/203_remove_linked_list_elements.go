package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// link to the task on leetcode
// https://leetcode.com/problems/remove-linked-list-elements/description/

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}
