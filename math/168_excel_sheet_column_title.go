package math

// link to the task on leetcode
// https://leetcode.com/problems/excel-sheet-column-title/description/

func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		columnNumber--
		ans = string('A'+rune(columnNumber%26)) + ans
		columnNumber /= 26
	}
	return ans
}
