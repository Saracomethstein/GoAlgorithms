package arrays_and_hashing

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	current := head
	ht := make(map[*ListNode]bool)
	for current != nil {
		if ht[current] == false {
			ht[current] = true
		} else {
			return true
		}
		current = current.Next
	}
	return false
}
