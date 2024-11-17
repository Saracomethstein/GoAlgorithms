package two_pointers

// link to the task on leetcode
// https://leetcode.com/problems/reverse-string/description/

func reverseString(s []byte) {
	lenS := len(s)
	for i := 0; i <= (lenS-2)/2; i++ {
		tmp := s[i]
		s[i] = s[lenS-1-i]
		s[lenS-1-i] = tmp
	}
}
