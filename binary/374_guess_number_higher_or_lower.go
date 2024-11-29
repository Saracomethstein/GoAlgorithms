package binary

// link to the task on leetcode
// https://leetcode.com/problems/guess-number-higher-or-lower/description/

func guessNumber(n int) int {
	start := 1
	end := n
	var mid int
	for i := 0; i < n; i++ {
		mid = (start + end) / 2
		if guess(mid) == -1 {
			end = mid - 1
		}

		if guess(mid) == 1 {
			start = mid + 1
		}

		if guess(mid) == 0 {
			break
		}
	}
	return mid
}

// don`t copy
func guess(n int) int {
	return 0
}
