package arrays_and_hashing

import "math"

// link to the task on leetcode
// https://leetcode.com/problems/container-with-most-water/description/

func maxArea(height []int) int {
	var lowH int = 0
	var maxH int = len(height) - 1
	var S float64
	for lowH < maxH {

		if S <= math.Min(float64(height[lowH]), float64(height[maxH]))*float64(maxH-lowH) {
			S = math.Min(float64(height[lowH]), float64(height[maxH])) * float64(maxH-lowH)
		}

		if height[lowH] < height[maxH] {
			lowH++
		} else if height[maxH] < height[lowH] {
			maxH--
		} else if height[maxH] == height[lowH] {
			lowH++
		}
	}
	return int(S)
}
