package stack

import "fmt"

// link to the task on leetcode
// https://leetcode.com/problems/largest-rectangle-in-histogram/description/
// I really don`t know how its worked

type StackValue struct {
	index  int
	height int
}

func LargestRectangleArea(heights []int) int {
	stack := []StackValue{}
	maxArea := 0
	var start int

	for i, h := range heights {
		start = i
		for len(stack) != 0 && stack[len(stack)-1].height > h {
			index, height := stack[len(stack)-1].index, stack[len(stack)-1].height
			stack = stack[0 : len(stack)-1]
			maxArea = Max(maxArea, height*(i-index))
			start = index
		}
		stack = append(stack, StackValue{start, h})
		fmt.Println(stack)
	}

	for _, h := range stack {
		maxArea = Max(maxArea, h.height*(len(heights)-h.index))
	}
	return maxArea
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
