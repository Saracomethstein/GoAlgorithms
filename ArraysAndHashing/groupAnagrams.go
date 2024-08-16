package main

import (
	"fmt"
)

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}

func groupAnagrams(strs []string) [][]string {
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
