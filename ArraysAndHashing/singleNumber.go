package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/single-number/description/

func SingleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}

	return single
}
