package arrays_and_hashing

import "math"

// link to the task on leetcode
// https://leetcode.com/problems/container-with-most-water/description/

func maxArea(height []int) int {
	var left int = 0
	var right int = len(height) - 1
	var S float64
	for left < right {

		tmpS := math.Min(float64(height[left]), float64(height[right])) * float64(right-left)
		if S <= tmpS {
			S = tmpS
		}

		if height[left] < height[right] {
			left++
		} else if height[right] < height[left] {
			right--
		} else if height[right] == height[left] {
			left++
		}
	}
	return int(S)
}
