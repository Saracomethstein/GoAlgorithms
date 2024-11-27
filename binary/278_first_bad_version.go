package binary

// link to the task on leetcode
// https://leetcode.com/problems/first-bad-version/description/

func firstBadVersion(n int) int {
	low := 1
	high := n

	for low < high {
		mid := low + (high-low)/2

		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// don`t copy
func isBadVersion(mid int) bool {
	return false
}
