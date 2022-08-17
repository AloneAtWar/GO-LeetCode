package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 20:21

 * @Note:	https://leetcode.cn/problems/course-schedule-ii/
			210. Course Schedule II.go
			210. 课程表 II
 **/

func findOrder(numCourses int, prerequisites [][]int) (result []int) {
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
			result = append(result, i)
		}
	}
	for len(list) != 0 {
		out := list[0]
		list = list[1:]
		for _, i := range graph[out] {
			inDegree[i]--
			if inDegree[i] == 0 {
				list = append(list, i)
				result = append(result, i)
			}
		}
	}
	if len(result) != numCourses {
		return nil
	}
	return
}
