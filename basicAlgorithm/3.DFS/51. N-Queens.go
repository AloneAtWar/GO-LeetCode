package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 1:51

 * @Note:	https://leetcode.cn/problems/n-queens/
			51. N-Queens
			51. N 皇后
 **/

func solveNQueens(n int) (resultView [][]string) {
	var dfs func(i int)
	result := make([]int, n)
	temp := "..................."
	column := make(map[int]bool, n)
	add := make(map[int]bool, n)
	cut := make(map[int]bool, n)
	//尝试第 i 行的 皇后应该放的位置
	dfs = func(i int) {
		if i == n {
			curr := make([]string, n)
			for j := 0; j < n; j++ {
				curr[j] = temp[:result[j]] + "Q" + temp[:n-result[j]-1]
			}
			resultView = append(resultView, curr)
			return
		}
		// n个位置都做一下尝试
		for j := 0; j < n; j++ {
			thisColumn, thisAdd, thisCut := j, j+i, j-i
			if !column[thisColumn] && !add[thisAdd] &&
				!cut[thisCut] {
				result[i] = j
				column[thisColumn], add[thisAdd], cut[thisCut] = true, true, true
				dfs(i + 1)
				column[thisColumn], add[thisAdd], cut[thisCut] = false, false, false
			}
		}
	}
	dfs(0)
	return
}
