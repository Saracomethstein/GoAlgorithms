package math

// link to the task on leetcode
// https://leetcode.com/problems/excel-sheet-column-number/description/

func titleToNumber(columnTitle string) int {
	mult := 1
	ans := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		ans += (int(columnTitle[i]) - int('A') + 1) * mult
		mult *= 26
	}
	return ans
}
