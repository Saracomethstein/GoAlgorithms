package arrays_and_hashing

import "math"

// link to the task on leetcode
// https://leetcode.com/problems/happy-number/description/

func isHappy(n int) bool {
	hashMap := make(map[float64]bool)
	ans := float64(n)
	for ans != 1.0 {
		ans = sumSlice(int(ans))

		if hashMap[float64(ans)] != true {
			hashMap[float64(ans)] = true
		} else {
			return false
		}
	}
	return true
}

func sumSlice(num int) float64 {
	sum := 0.0
	for num > 0 {
		digit := num % 10
		num /= 10

		sum += math.Pow(float64(digit), 2)
	}
	return sum
}
