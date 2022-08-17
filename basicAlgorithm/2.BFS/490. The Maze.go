package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 20:25

 * @Note:	https://leetcode.cn/problems/the-maze/
			490. The Maze
			490. 迷宫
 **/

func hasPath(maze [][]int, start []int, destination []int) bool {
	n, m := len(maze), len(maze[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	visited[start[0]][start[1]] = true
	paths := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	list := [][]int{start}
	for len(list) != 0 {
		out := list[0]
		if out[0] == destination[0] && out[1] == destination[1] {
			return true
		}
		list = list[1:]
		for _, path := range paths {
			newX, newY := out[0], out[1]
			for {
				newX, newY = newX+path[0], newY+path[1]
				if newX < 0 || newY < 0 || newX >= n || newY >= m || maze[newX][newY] == 1 {
					break
				}
			}
			newX -= path[0]
			newY -= path[1]
			if visited[newX][newY] {
				continue
			}
			list = append(list, []int{newX, newY})
			visited[newX][newY] = true
		}
	}
	return false
}
