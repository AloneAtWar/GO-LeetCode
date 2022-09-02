package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 16:36

 * @Note:	https://leetcode.cn/problems/max-area-of-island/
			695. Max Area of Island
			695. 岛屿的最大面积
 **/

func maxAreaOfIsland(grid [][]int) (result int) {
	n, m := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int) int
	dfs = func(i, j int) (result int) {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			result += dfs(newI, newJ)
		}
		return result + 1
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
