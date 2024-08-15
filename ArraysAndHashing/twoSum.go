package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 0
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var hash = make(map[int]int)
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		}
		hash[num] = i
	}
	return nil
}
