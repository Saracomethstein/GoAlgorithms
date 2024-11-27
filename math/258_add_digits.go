package math

// link to the task on leetcode
// https://leetcode.com/problems/add-digits/description/

func addDigits(num int) int {
	current := num
	for current >= 10 {
		current = sumDigits(current)
	}

	return current
}

func sumDigits(num int) int {
	current := num
	sum := 0

	for current > 0 {
		digit := current % 10
		current = current / 10
		sum += digit
	}
	return sum
}
