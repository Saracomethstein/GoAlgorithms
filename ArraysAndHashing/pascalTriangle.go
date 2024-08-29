package arrays_and_hashing

import "fmt"

// link to the task on leetcode
// https://leetcode.com/problems/pascals-triangle/description/

func Generate(numRows int) [][]int {
	res := make([][]int, 0)

	// generate helper matrix
	array := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		array[i] = make([]int, numRows)
	}

	// fill 1 to edges of matrix
	for i := 0; i < numRows; i++ {
		array[0][i], array[i][0] = 1, 1
	}

	// calculate value
	for i := 1; i < numRows; i++ {
		for j := 1; j < numRows-i; j++ {
			array[i][j] = array[i-1][j] + array[i][j-1]
		}
	}

	fmt.Println(array)

	// build result
	for i := 0; i < numRows; i++ {
		temp := make([]int, 0)
		for j := 0; j <= i; j++ {
			temp = append(temp, array[i-j][j])
		}
		res = append(res, temp)
	}
	return res
}
