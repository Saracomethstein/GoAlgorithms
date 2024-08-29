package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(arrays_and_hashing.ContainsNearbyDuplicate(nums, 2))
}
