package arrays_and_hashing

import (
	"sort"
)

// link to the task on leetcode
// https://leetcode.com/problems/3sum/description/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	hashMap := make(map[[3]int]bool)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				tmpArr := [3]int{nums[i], nums[left], nums[right]}

				if !hashMap[tmpArr] {
					hashMap[tmpArr] = true
					result = append(result, tmp)
				}
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
