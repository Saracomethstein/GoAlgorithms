package stack

import "fmt"

func CarFleet(target int, position []int, speed []int) int {
	stack := make([]int, 0)
	for i := len(position) - 1; i >= 0; i-- {
		checkTarget := position[i] + speed[i]

		if checkTarget <= target {
			if stack[len(stack)-1] != checkTarget {
				stack = append(stack, checkTarget)
			}
		}

		fmt.Println(stack)

	}
	return len(stack)
}
