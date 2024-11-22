package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/intersection-of-two-arrays/description/

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]bool)
	result := []int{}
	for _, v := range nums1 {
		hashMap[v] = false
	}

	for _, v := range nums2 {
		if _, ok := hashMap[v]; ok && hashMap[v] != true {
			hashMap[v] = true
			result = append(result, v)
		}
	}

	return result
}
