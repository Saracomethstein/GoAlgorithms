package arrays_and_hashing

import (
	"strings"
)

// link to the task on leetcode
// https://leetcode.com/problems/word-pattern/description/

func wordPattern(pattern string, s string) bool {
	hashMap := make(map[string]string)
	reverseMap := make(map[string]string)
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		letter := string(pattern[i])
		word := words[i]

		if mappedWord, ok := hashMap[letter]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			if existingLetter, ok := reverseMap[word]; ok && existingLetter != letter {
				return false
			}
			hashMap[letter] = word
			reverseMap[word] = letter
		}
	}

	return true
}
