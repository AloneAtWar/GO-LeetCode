package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 15:18

 * @Note:	https://leetcode.cn/problems/course-schedule/
			207. Course Schedule
			207. 课程表
 **/

// 老题 之前学BFS的时候 使用的 这次就换成三色法的 DFS把
func canFinish(numCourses int, prerequisites [][]int) bool {
	valid := true
	state := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[0]] = append(graph[prerequisite[0]], prerequisite[1])
	}
	//0 为搜索 1为正在搜索 2 搜索结束
	var dfs func(i int)
	dfs = func(i int) {
		if !valid {
			return
		}
		state[i] = 1
		for _, v := range graph[i] {
			if state[v] == 0 {
				dfs(v)
			} else if state[v] == 1 {
				valid = false
				return
			}
		}
		state[i] = 2
	}
	for i := 0; i < numCourses; i++ {
		if !valid {
			break
		}
		if state[i] == 0 {
			dfs(i)
		}
	}

	return valid
}
