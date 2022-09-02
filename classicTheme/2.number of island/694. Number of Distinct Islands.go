package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 18:02

 * @Note:	https://leetcode.cn/problems/number-of-distinct-islands/
			694. Number of Distinct Islands
			694. 不同岛屿的数量
 **/

func numDistinctIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 左是 1 右是2 上是3 下是4
	dirs := [][]int{{0, -1, 1}, {0, 1, 2}, {1, 0, 3}, {-1, 0, 4}}
	table := make(map[string]bool)
	var dfs func(i, j int, curr string, next int) string
	dfs = func(i, j int, curr string, byte int) string {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return curr
		}
		grid[i][j] = 0
		curr = fmt.Sprintf("%s%d", curr, byte)
		for _, dir := range dirs {
			newX, newY := i+dir[0], j+dir[1]
			curr = dfs(newX, newY, curr, dir[2])
		}
		// 加个结束符号  ！！！！
		curr = fmt.Sprintf("%s%d", curr, 0)
		return curr
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				s := dfs(i, j, "", 0)
				table[s] = true
			}
		}
	}
	return len(table)
}
