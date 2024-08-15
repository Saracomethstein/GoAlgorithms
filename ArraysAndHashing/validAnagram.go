package main

import (
	"fmt"
	"sort"
)

func main() {
	sA := "anagram"
	sB := "nagaram"
	fmt.Println(isAnagram(sA, sB))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sliceA := []byte(s)
	sliceB := []byte(t)

	sort.Slice(sliceA, func(i, j int) bool { return sliceA[i] < sliceA[j] })
	sort.Slice(sliceB, func(i, j int) bool { return sliceB[i] < sliceB[j] })

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}
