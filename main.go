package main

import (
	stack "GoAlgorithms/Stack"
	"fmt"
)

func main() {
	position := []int{6, 2, 17}
	speed := []int{3, 9, 2}
	fmt.Println("Answer:", stack.CarFleet(20, position, speed))
}
