package math

// link to the task on leetcode
// https://leetcode.com/problems/power-of-four/description/

func isPowerOfFour(n int) bool {
	return (n&(n-1)) == 0 && (n-1)%3 == 0
}
