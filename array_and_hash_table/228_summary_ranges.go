package arrays_and_hashing

import "strconv"

// link to the task on leetcode
// https://leetcode.com/problems/summary-ranges/description/

func summaryRanges(nums []int) []string {
	if nums == nil || len(nums) == 0 {
		return []string{}
	}

	l := len(nums)
	var ans []string

	for i := 0; i < l; i++ {
		num := nums[i]

		for i < l-1 && nums[i+1] == nums[i]+1 {
			i++
		}

		if num == nums[i] {
			ans = append(ans, strconv.Itoa(num))
			continue
		}
		ans = append(ans, strconv.Itoa(num)+"->"+strconv.Itoa(nums[i]))
	}
	return ans
}
