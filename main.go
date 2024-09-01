package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 1}
	nums1 := []int{2, 2}
	fmt.Println(arrays_and_hashing.Intersect(nums, nums1))
}
