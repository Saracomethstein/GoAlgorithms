package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func Intersect(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	result := []int{}
	for _, v := range nums1 {
		hashMap[v] += 1
	}

	for _, v := range nums2 {
		if _, ok := hashMap[v]; ok && hashMap[v] > 0 {
			hashMap[v] -= 1
			result = append(result, v)
		}
	}

	return result
}
