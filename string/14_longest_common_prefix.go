package string

// link to the task on leetcode
// https://leetcode.com/problems/longest-common-prefix/description/

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return prefix
			}
		}
	}
	return prefix
}
