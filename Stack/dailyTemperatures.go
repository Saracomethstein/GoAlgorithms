package stack

// link to the task on leetcode
// https://leetcode.com/problems/daily-temperatures/submissions/1368818322/

func DailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		j := i + 1
		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			if answer[j] <= 0 {
				break
			}
			j += answer[j]
		}

		if j < len(temperatures) && temperatures[j] > temperatures[i] {
			answer[i] = j - i
		}
	}
	return answer
}
