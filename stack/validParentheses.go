package stack

// link to the task on leetcode
// https://leetcode.com/problems/valid-parentheses/submissions/1367843292/

func IsValid(s string) bool {
	hashMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0)

	for _, c := range []rune(s) {
		pair, ok := hashMap[c]
		if !ok {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != pair {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
