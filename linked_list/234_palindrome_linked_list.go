package linked_list

// link to the task on leetcode
// https://leetcode.com/problems/palindrome-linked-list/description/

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	stack := []*ListNode{}
	current := head

	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	j := 0
	for i := len(stack) - 1; i >= len(stack)/2; i-- {
		if stack[i].Val != stack[j].Val {
			return false
		}
		j++
	}
	return true
}
