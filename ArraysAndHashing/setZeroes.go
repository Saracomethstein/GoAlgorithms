package arrays_and_hashing

// link to the task on leetcode 
// https://leetcode.com/problems/set-matrix-zeroes/description/

func SetZeros(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				col[j] = 1
				row[i] = 1
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if col[j] == 1 || row[i] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
