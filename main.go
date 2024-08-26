package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	heights := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	fmt.Println("Answer:", stack.LargestRectangleArea(heights))
}
