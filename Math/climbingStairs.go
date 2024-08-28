package math

// link to the task on leetcode
// https://leetcode.com/problems/climbing-stairs/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	n1 := 1
	n2 := 2

	for i := 3; i <= n; i++ {
		next := n1
		n1 = n2
		n2 = next + n2
	}
	return n2
}
