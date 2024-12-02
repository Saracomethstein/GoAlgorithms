package math

import (
	"math"
)

// link to the task on leetcode
// https://leetcode.com/problems/valid-perfect-square/description/

func isPerfectSquare(num int) bool {
	x := float64(num)
	z := float64(num) / 2
	s := 0.0

	for i := 0; i < 20; i++ {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-s) < 1e-10 {
			break
		}
		s = z
	}

	return math.Abs(z-math.Round(z)) < 1e-10
}
