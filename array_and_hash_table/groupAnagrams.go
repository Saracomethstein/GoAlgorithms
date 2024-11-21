package arrays_and_hashing

// link to the task on leetcode 
// https://leetcode.com/problems/group-anagrams/description/

func GroupAnagrams(strs []string) [][]string {
	var anagrams = make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int

		for _, c := range s {
			count[c-'a']++
		}
		anagrams[count] = append(anagrams[count], s)
	}

	var result = make([][]string, len(anagrams))

	i := 0
	for _, c := range anagrams {
		result[i] = c
		i++
	}

	return result
}
