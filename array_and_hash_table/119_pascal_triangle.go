package arrays_and_hashing

// link to the task on leetcode
// https://leetcode.com/problems/pascals-triangle-ii/description/

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		res := []int{1}
		return res
	}
	rowIndex++

	res := make([][]int, 0)

	array := make([][]int, rowIndex)
	for i := 0; i < rowIndex; i++ {
		array[i] = make([]int, rowIndex)
	}

	for i := 0; i < rowIndex; i++ {
		array[0][i], array[i][0] = 1, 1
	}

	for i := 1; i < rowIndex; i++ {
		for j := 1; j < rowIndex-i; j++ {
			array[i][j] = array[i-1][j] + array[i][j-1]
		}
	}

	for i := 0; i < rowIndex; i++ {
		temp := make([]int, 0)
		for j := 0; j <= i; j++ {
			temp = append(temp, array[i-j][j])
		}
		res = append(res, temp)
	}

	return res[rowIndex-1]
}
