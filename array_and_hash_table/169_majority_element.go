package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/majority-element/description/

func MajorityElement(nums []int) int {
	hashMap := make(map[int]int)
	count := 0
	element := 0

	for _, n := range nums {
		hashMap[n]++
	}

	for k, v := range hashMap {
		if v > count {
			count = v
			element = k
		}
	}

	return element
}
