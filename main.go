package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	//nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(arrays_and_hashing.LongestConsecutive(nums))
}
