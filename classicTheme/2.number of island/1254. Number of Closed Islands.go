package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 18:26

 * @Note:	https://leetcode.cn/problems/number-of-closed-islands/
			1254. Number of Closed Islands
			1254. 统计封闭岛屿的数目
 **/

func closedIsland(grid [][]int) (result int) {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) bool
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dfs = func(i, j int) bool {
		if i < 0 || j < 0 || i >= m || j >= n {
			return false
		}
		if grid[i][j] == 1 {
			return true
		}
		grid[i][j] = 1
		result := true
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			if !dfs(newI, newJ) {
				result = false
			}
		}
		return result
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && dfs(i, j) {
				result++
			}
		}
	}
	return
}
