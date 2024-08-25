package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	s := "([()])"
	fmt.Println(stack.IsValid(s))
}
