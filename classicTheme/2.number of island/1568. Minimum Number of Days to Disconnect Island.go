package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 19:32

 * @Note:	https://leetcode.cn/problems/minimum-number-of-days-to-disconnect-island/
			1568. Minimum Number of Days to Disconnect Island
			1568. 使陆地分离的最少天数
 **/

func minDays(grid [][]int) int {
	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{-1, 0}, []int{1, 0}}
	graph := map[[2]int][][2]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				if graph[[2]int{i, j}] == nil {
					graph[[2]int{i, j}] = [][2]int{}
				}
				for _, d := range dirs {
					x, y := i+d[0], j+d[1]
					if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] == 1 {
						graph[[2]int{i, j}] = append(graph[[2]int{i, j}], [2]int{x, y})
					}
				}
			}
		}
	}
	// 特殊情况处理
	if len(graph) == 1 {
		return 1
	}
	if len(graph) == 0 {
		return 0
	}
	dfn, low := map[[2]int]int{}, map[[2]int]int{}
	for u := range graph {
		dfn[u] = 0
		low[u] = 0
	}
	ts := 1
	cuts := map[[2]int]bool{} // 求割点
	var trajan func(u, parent [2]int)
	trajan = func(u, parent [2]int) {
		dfn[u], low[u] = ts, ts
		ts++
		children := 0
		//for v := range graph[u]{ 注意 这个地方 索引 和  值 经常搞混
		for _, v := range graph[u] {
			if v == parent {
				continue
			}
			if dfn[v] != 0 { // 处理回边
				low[u] = min(low[u], dfn[v])
			} else {
				children++
				trajan(v, u)
				low[u] = min(low[u], low[v])
				// 判断 割点
				if parent == [2]int{-1, -1} && children >= 2 {
					cuts[u] = true
				} else if parent != [2]int{-1, -1} && low[v] >= dfn[u] {
					cuts[u] = true
				}
			}
		}
	}
	// trajan 算法
	cnt := 0
	for u := range graph {
		if dfn[u] != 0 {
			continue
		}
		cnt++
		if cnt > 1 {
			return 0
		}
		trajan(u, [2]int{-1, -1})
	}
	if len(cuts) > 0 {
		return 1
	}
	return 2
}
func min(nums ...int) int {
	m := nums[0]
	for _, c := range nums {
		if m > c {
			m = c
		}
	}
	return m
}
