package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	var sub []byte
	var res int

	for i := range s {
		if check, _ := contains(sub, s[i]); !check {
			sub = append(sub, s[i])
			res = max(res, len(sub))
		} else {
			check, index := contains(sub, s[i])
			if check {
				sub = append(sub, s[i])
				sub = sub[index+1:]
			}
		}
	}

	return res
}

func contains(s []byte, e byte) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
