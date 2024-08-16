package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	var frequentM = make(map[int]int)
	var frequentS = make([][]int, len(nums)+1)
	var res = make([]int, 0)

	for _, num := range nums {
		if count, ok := frequentM[num]; ok {
			frequentM[num] = count + 1
		} else {
			frequentM[num] = 1
		}
	}

	for num, count := range frequentM {
		frequentS[count] = append(frequentS[count], num)
	}

	for i := len(frequentS) - 1; i >= 0; i-- {
		res = append(res, frequentS[i]...)
		if len(res) == k {
			return res
		}
	}

	return res
}
