package arrays_and_hashing

import (
	"fmt"
	"math"
)

// link to the task on leetcode
// https://leetcode.com/problems/contains-duplicate-ii/description/

func ContainsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int)

	for i, v := range nums {
		if val, ok := hashMap[v]; ok && math.Abs(float64(i)-float64(val)) <= float64(k) {
			return true
		}

		hashMap[v] = i
		fmt.Println(hashMap)
	}

	return false
}
