package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/isomorphic-strings/description/

func isIsomorphic(s string, t string) bool {
	hashMap := make(map[string]string)
	reverseMap := make(map[string]string)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		letterS := string(s[i])
		letterT := string(t[i])

		if mappedLetter, ok := hashMap[letterS]; ok {
			if mappedLetter != letterT {
				return false
			}
		} else {
			if existingLetter, ok := reverseMap[letterT]; ok && existingLetter != letterS {
				return false
			}
			hashMap[letterS] = letterT
			reverseMap[letterT] = letterS
		}
	}
	return true
}
