package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	arrays_and_hashing.Merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
