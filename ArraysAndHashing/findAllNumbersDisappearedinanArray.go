package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

func FindDisappearedNumbers(nums []int) []int {
	hashMap := make(map[int]bool)
	result := []int{}
	for i := 1; i <= len(nums); i++ {
		hashMap[i] = false
	}

	for _, v := range nums {
		if _, ok := hashMap[v]; ok {
			hashMap[v] = true
		}
	}

	for v := range hashMap {
		if hashMap[v] == false {
			result = append(result, v)
		}
	}

	return result
}
