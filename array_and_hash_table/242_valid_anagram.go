package arrays_and_hashing

import "sort"

// // link to the task on leetcode
// https://leetcode.com/problems/valid-anagram/description/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sliceA := []byte(s)
	sliceB := []byte(t)

	sort.Slice(sliceA, func(i, j int) bool { return sliceA[i] < sliceA[j] })
	sort.Slice(sliceB, func(i, j int) bool { return sliceB[i] < sliceB[j] })

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
