package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 1:20

 * @Note:	https://leetcode.cn/problems/sudoku-solver/
			37. Sudoku Solver
			37. 解数独
 **/

func solveSudoku(board [][]byte) {
	row := make([]map[int]bool, 9)
	column := make([]map[int]bool, 9)
	square := make([]map[int]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[int]bool, 9)
		column[i] = make(map[int]bool, 9)
		square[i] = make(map[int]bool, 9)
	}
	for i, bytes := range board {
		for j, v := range bytes {
			if v != '.' {
				row[i][int(v-'0')] = true
				column[j][int(v-'0')] = true
				square[check(i, j)][int(v-'0')] = true
			}
		}
	}
	var dfs func() bool
	dfs = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					squareNum := check(i, j)
					for m := 1; m <= 9; m++ {
						if !row[i][m] && !column[j][m] && !square[squareNum][m] {
							row[i][m] = true
							column[j][m] = true
							square[check(i, j)][m] = true
							board[i][j] = byte('0' + m)
							if dfs() {
								return true
							} else {
								row[i][m] = false
								column[j][m] = false
								square[squareNum][m] = false
								board[i][j] = '.'
							}
						}
					}
					return false
				}
			}
		}
		return true
	}
	dfs()
	return
}

// 用来算他是第几个9x9方格里头的 从左到有 从上到下 1 2 3 4 5	....
//[["5","3","0","2","7","6","4","1","8"],
//	["6","2","4","1","9","5","3","0","7"],
//	["1","9","8","3","4","0","5","6","2"],["8","1","2","7","6","4","0","5","3"],["4","0","6","8","5","3","7","2","1"],["7","5","3","0","2","1","8","4","6"],["0","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","6","8","2","1","7","9"]]
func check(x, y int) int {
	return (y/3)*3 + (x / 3)
}
