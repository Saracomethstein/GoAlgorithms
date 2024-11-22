package stack

import (
	"strconv"
)

// link to the task on leetcode
// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
	result, _ := strconv.Atoi(tokens[0])
	stack := make([]string, 0)

	for _, t := range tokens {
		if t != "+" && t != "-" && t != "*" && t != "/" {
			stack = append(stack, t)
		} else {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			result = parse(a, b, t)
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = strconv.Itoa(result)
		}
	}

	return result
}

func parse(a, b, operator string) int {
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)

	if operator == "+" {
		return intA + intB
	} else if operator == "-" {
		return intA - intB
	} else if operator == "*" {
		return intA * intB
	} else if operator == "/" {
		return intA / intB
	} else {
		return 0
	}
}
