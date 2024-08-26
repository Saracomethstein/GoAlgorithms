package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	tokens := []string{"4"}
	fmt.Println(stack.EvalRPN(tokens))
}
