package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/find-the-difference/description/

func findTheDifference(s string, t string) byte {
	var sumS, sumT byte

	for _, v := range t {
		sumT += byte(v)
	}

	for _, v := range s {
		sumS += byte(v)
	}

	return sumT - sumS
}
