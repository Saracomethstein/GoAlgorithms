package main

import (
	arrays_and_hashing "GoAlgorithms/ArraysAndHashing"
	"fmt"
)

func main() {
	matrix := [][]int{{0, 0, 0, 5}, {4, 3, 1, 4}, {0, 1, 1, 4}, {1, 2, 1, 3}, {0, 0, 1, 1}}
	arrays_and_hashing.SetZeros(matrix)
	fmt.Println(matrix)
}
