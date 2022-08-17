package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 20:10

 * @Note:	https://leetcode.cn/problems/course-schedule/
			207. Course Schedule
			207. 课程表
 **/

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	outDegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		outDegree[prerequisite[1]]++
		inDegree[prerequisite[0]]++
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	var list []int
	for i, v := range inDegree {
		if v == 0 {
			list = append(list, i)
		}
	}
	size := 0
	for len(list) != 0 {
		out := list[0]
		list = list[1:]
		size++
		for _, i := range graph[out] {
			inDegree[i]--
			if inDegree[i] == 0 {
				list = append(list, i)
			}
		}
	}
	return size == numCourses
}
