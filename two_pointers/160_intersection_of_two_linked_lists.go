package two_pointers

// link to the task on leetcode
// https://leetcode.com/problems/intersection-of-two-linked-lists/description/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	slow, fast := headA, headB
	for slow != fast {
		if slow == nil {
			slow = headB
		} else {
			slow = slow.Next
		}

		if fast == nil {
			fast = headA
		} else {
			fast = fast.Next
		}
	}
	return slow
}
