package arrays_and_hashing

import (
	"math"
)

func MaxArea(height []int) int {
	var lenH int = len(height) - 1
	var lowH int = 0
	var maxH int = lenH
	var S float64
	for {

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

		if maxH == lowH {
			break
		}
	}
	return int(S)
}
