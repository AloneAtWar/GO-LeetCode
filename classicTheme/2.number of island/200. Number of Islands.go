package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 16:27

 * @Note:	https://leetcode.cn/problems/number-of-islands/
			200. Number of Islands
			200. 岛屿数量
 **/

func numIslands(grid [][]byte) (result int) {
	n, m := len(grid), len(grid[0])
	visited := make([][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, m)
	}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == '0' || visited[i][j] == 1 {
			return
		}
		visited[i][j] = 1
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			dfs(newI, newJ)
		}
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 0 && grid[i][j] == '1' {
				result++
				dfs(i, j)
			}
		}
	}
	return
}
