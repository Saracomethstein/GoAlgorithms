package math

// link to the task on leetcode
// https://leetcode.com/problems/power-of-two/description/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	if n&(n-1) == 0 {
		return true
	}
	return false
}
