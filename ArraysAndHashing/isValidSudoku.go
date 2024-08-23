package arrays_and_hashing

// link to the task on leetcode 
// https://leetcode.com/problems/valid-sudoku/description/

func IsValidSudoku(board [][]byte) bool {
	mapBoard := make(map[string]bool)
	for i := range board {
		for j := range board {
			val := string(board[i][j])

			if val == "." {
				continue
			}

			_, ok1 := mapBoard[val+"row"+string(i)]
			_, ok2 := mapBoard[val+"column"+string(j)]
			_, ok3 := mapBoard[val+"grid"+string(i/3)+"-"+string(j/3)]

			if ok1 || ok2 || ok3 {
				return false
			} else {
				mapBoard[val+"row"+string(i)] = true
				mapBoard[val+"column"+string(j)] = true
				mapBoard[val+"grid"+string(i/3)+"-"+string(j/3)] = true
			}
		}
	}
	return true
}
