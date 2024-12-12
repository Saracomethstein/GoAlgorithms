package two_pointers

// link to the task on leetcode
// https://leetcode.com/problems/longest-palindromic-substring/description/

func LongestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	start, maxLength := 0, 1

	for i := 0; i < n; i++ {
		expandAroundCenter(s, i, i, &start, &maxLength)
		expandAroundCenter(s, i, i+1, &start, &maxLength)
	}

	return s[start : start+maxLength]
}

func expandAroundCenter(s string, left, right int, start, maxLength *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		if right-left+1 > *maxLength {
			*start = left
			*maxLength = right - left + 1
		}
		left--
		right++
	}
}
