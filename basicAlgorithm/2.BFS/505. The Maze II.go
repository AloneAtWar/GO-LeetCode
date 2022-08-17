package main

import "container/heap"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 20:37

 * @Note:	https://leetcode.cn/problems/the-maze-ii/
			505. The Maze II.go
505. 迷宫 II
 **/

type Point [][]int

func (m *Point) Len() int {
	return len(*m)
}

func (m *Point) Less(i, j int) bool {
	return (*m)[i][2] < (*m)[j][2]
}

func (m *Point) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *Point) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *Point) Pop() interface{} {
	out := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return out
}
func shortestDistance(maze [][]int, start []int, destination []int) int {
	n, m := len(maze), len(maze[0])
	visited := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make(map[int]int, m)
	}
	visited[start[0]][start[1]] = 0
	paths := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	h := &Point{append(start, 0)}
	for h.Len() != 0 {
		out := heap.Pop(h).([]int)
		if out[0] == destination[0] && out[1] == destination[1] {
			return out[2]
		}
		if visited[out[0]][out[1]] < out[2] {
			continue
		}
		for _, path := range paths {
			newX, newY := out[0], out[1]
			p := 0
			for {
				newX, newY = newX+path[0], newY+path[1]
				if newX < 0 || newY < 0 || newX >= n || newY >= m || maze[newX][newY] == 1 {
					break
				}
				p++
			}
			newX -= path[0]
			newY -= path[1]
			if v, ok := visited[newX][newY]; !ok || p+out[2] < v {
				visited[newX][newY] = p + out[2]
				heap.Push(h, []int{newX, newY, p + out[2]})
			}
		}
	}
	return -1
}
