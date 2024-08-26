package stack

import (
	"fmt"
	"sort"
)

// link to the task on leetcode
// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

func LargestRectangleArea(heights []int) int {
	lowP := 0
	maxP := len(heights) - 1
	var S int

	fmt.Println("S:\tLowP:\tMaxP:\tForm:\tMinHeight:\t")

	for {
		if S < getMinHeight(heights[lowP:maxP+1])*((maxP+1)-lowP) {
			S = getMinHeight(heights[lowP:maxP+1]) * ((maxP + 1) - lowP)
		}

		if maxP == lowP {
			break
		}

		fmt.Printf("%d\t%d\t%d\t%d\t%d\t\n", S, lowP, maxP, (maxP+1)-lowP, getMinHeight(heights[lowP:maxP]))

		if heights[lowP] < heights[maxP] {
			lowP++
		} else if heights[maxP] < heights[lowP] {
			maxP--
		} else {
			lowP++
		}
	}

	return S
}

func getMinHeight(slice []int) int {
	fmt.Println(slice)
	temp := make([]int, len(slice))
	copy(temp, slice)
	sort.Ints(temp)
	return temp[0]
}
