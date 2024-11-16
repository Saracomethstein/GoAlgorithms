package two_pointers

// link to the task on leetcode
// https://leetcode.com/problems/reverse-vowels-of-a-string/description/

func reverseVowels(s string) string {
	vowels := map[rune]struct {
	}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}

	word := []byte(s)
	start, end := 0, len(word)-1

	for start < end {
		for start < end && !isVowel(rune(word[start]), vowels) {
			start++
		}

		for start < end && !isVowel(rune(word[end]), vowels) {
			end--
		}

		word[start], word[end] = word[end], word[start]
		start++
		end--
	}

	return string(word)
}

func isVowel(c rune, vowels map[rune]struct{}) bool {
	_, exists := vowels[c]
	return exists
}
