package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 3, 2, 13}
	arrays_and_hashing.MoveZeroes(nums)
	fmt.Println(nums)
}
