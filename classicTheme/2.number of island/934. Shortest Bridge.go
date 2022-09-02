package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 18:57

 * @Note:	https://leetcode.cn/problems/shortest-bridge/
			934. Shortest Bridge
			934. 最短的桥
 **/

// DFS+双向 BFS 真的这个代码写的是 真的臭！
func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	list1, list2 := [][]int{}, [][]int{}
	table1, table2 := map[int]bool{}, map[int]bool{}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dfs func(i, j int, replace int, table map[int]bool, list *[][]int)
	dfs = func(i, j int, replace int, table map[int]bool, list *[][]int) {
		grid[i][j] = replace
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			if newI < 0 || newJ < 0 || newI >= m || newJ >= n || grid[newI][newJ] != 1 {
				table[i*n+j] = true
			} else {
				dfs(newI, newJ, replace, table, list)
			}
		}
		if table[i*n+j] {
			*list = append(*list, []int{i, j})
		}
	}
	replce := 2
	table := table1
	list := &list1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j, replce, table, list)
				if replce == 2 {
					replce = 3
					table = table2
					list = &list2
				}
			}
		}
	}
	step := 0
	for len(list1) != 0 {
		temp := [][]int{}

		for _, ints := range list1 {
			i, j := ints[0], ints[1]
			for _, dir := range dirs {
				newI, newJ := i+dir[0], j+dir[1]
				if newI < 0 || newJ < 0 || newI >= m || newJ >= n || table1[newI*n+newJ] {
					continue
				}

				if table2[newI*n+newJ] {
					return step
				}
				if grid[newI][newJ] != 0 {
					continue
				}
				table1[newI*n+newJ] = true
				temp = append(temp, []int{newI, newJ})
			}
		}
		if len(temp) > len(list2) {
			list1, list2 = list2, temp
			table1, table2 = table2, table1
		} else {
			list1 = temp
		}
		step++
	}
	return -1
}
