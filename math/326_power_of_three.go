package math

// link to the task on leetcode
// https://leetcode.com/problems/power-of-three/description/

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
