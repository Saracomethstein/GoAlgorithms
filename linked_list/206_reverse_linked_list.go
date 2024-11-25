package linked_list

// link to the task on leetcode
// https://leetcode.com/problems/reverse-linked-list/description/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
