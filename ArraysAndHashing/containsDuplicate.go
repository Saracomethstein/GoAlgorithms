package main

import (
	"fmt"
)

func main() {
	numsA := []int{1, 2, 3, 1}
	numsB := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(numsA))
	fmt.Println(containsDuplicate(numsB))
}

func containsDuplicate(nums []int) bool {
	var noneDup = make(map[int]bool)

	for _, v := range nums {
		if _, ok := noneDup[v]; ok {
			return true
		}
		noneDup[v] = true
	}

	return false
}
