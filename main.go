package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	fmt.Println(stack.CarFleet(12, position, speed))
}
