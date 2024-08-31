package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/missing-number/description/

func MissingNumber(nums []int) int {
	sumNums := 0
	needSum := 0

	j := 1
	for i := range nums {
		sumNums += nums[i]
		needSum += j
		j++
	}

	return needSum - sumNums
}
