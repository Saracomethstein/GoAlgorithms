package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	height := []int{2, 1, 5, 6, 2, 3}
	fmt.Println("Answer:", stack.LargestRectangleArea(height))
}
