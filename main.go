package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 4
	fmt.Println(arrays_and_hashing.TwoSum(nums, target))
}
