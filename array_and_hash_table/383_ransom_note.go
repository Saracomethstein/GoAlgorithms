package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/ransom-note/description/

func canConstruct(ransomNote string, magazine string) bool {
	hashMap := make(map[rune]int)
	for _, char := range magazine {
		hashMap[char]++
	}

	for _, char := range ransomNote {
		if hashMap[char] <= 0 {
			return false
		}
		hashMap[char]--
	}

	return true
}
