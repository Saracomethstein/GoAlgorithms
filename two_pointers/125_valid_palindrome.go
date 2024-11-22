package two_pointers

import (
	"strings"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !(s[i] >= 97 && s[i] <= 122 || s[i] >= 48 && s[i] <= 57 || s[i] >= 65 && s[i] <= 90) {
			i++
		}

		for i < j && !(s[j] >= 97 && s[j] <= 122 || s[j] >= 48 && s[j] <= 57 || s[j] >= 65 && s[j] <= 90) {
			j--
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}
	return true
}
