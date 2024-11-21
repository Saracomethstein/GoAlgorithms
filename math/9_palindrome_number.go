package math

// link to the task on leetcode
// https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	var y int = x
	var res int

	if y > 0 {
		for {
			res += y % 10
			y = y / 10

			if y > 0 {
				res *= 10
			} else {
				break
			}
		}

		if x == res {
			return true
		}
	} else if y == 0 {
		return true
	}

	return false
}
