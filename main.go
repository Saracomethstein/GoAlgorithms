package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	tokens := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(tokens)
	fmt.Println(stack.DailyTemperatures(tokens))
}
