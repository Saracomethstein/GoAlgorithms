package arrays_and_hashing

import "math"

// link to the task on leetcode
// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		minRightX := math.MaxInt64
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		minRightY := math.MaxInt64
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return (float64(math.Max(float64(maxLeftX), float64(maxLeftY))) +
					float64(math.Min(float64(minRightX), float64(minRightY)))) / 2.0
			} else {
				return math.Max(float64(maxLeftX), float64(maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}
