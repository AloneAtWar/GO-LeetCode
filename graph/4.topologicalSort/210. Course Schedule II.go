package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 15:24

 * @Note:

 **/

func findOrder(numCourses int, prerequisites [][]int) (result []int) {
	valid := true
	state := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		// 注意此处边的顺序是反序
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
		result = append(result, i)
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
	if !valid {
		return []int{}
	} else {
		return flip(result)
	}
}

func flip(list []int) []int {
	left, right := 0, len(list)-1
	for left < right {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	return list
}
