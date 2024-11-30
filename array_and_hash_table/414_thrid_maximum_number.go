package arrays_and_hashing

import "sort"

// link to the task on leetcode
// https://leetcode.com/problems/third-maximum-number/description/

func thirdMax(nums []int) int {
	unique := make(map[int]struct{})

	for _, num := range nums {
		unique[num] = struct{}{}
	}

	var uniqueNums []int
	for num := range unique {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(uniqueNums)))

	if len(uniqueNums) < 3 {
		return uniqueNums[0]
	}

	return uniqueNums[2]
}
