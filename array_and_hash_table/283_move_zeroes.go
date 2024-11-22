package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/move-zeroes/description/

func moveZeroes(nums []int) {
	temp := make([]int, len(nums))
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			temp[j] = nums[i]
			j++
		}
	}
	copy(nums, temp)
}
