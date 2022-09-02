package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 16:41

 * @Note:	https://leetcode.cn/problems/island-perimeter/
			463. Island Perimeter
			463. 岛屿的周长
 **/

func islandPerimeter(grid [][]int) (result int) {
	n, m := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
			return 1
		}
		if grid[i][j] == -1 {
			return 0
		}
		grid[i][j] = -1
		return dfs(i+1, j) + dfs(i-1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				this := dfs(i, j)
				if this > result {
					result = this
				}
			}
		}
	}
	return
}
