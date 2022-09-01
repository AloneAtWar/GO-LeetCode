package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 18:10

 * @Note:   https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/
			2127. Maximum Employees to Be Invited to a Meeting
			2127. 参加会议的最多员工数
 **/

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	graph := make([][]int, n)
	for i, v := range favorite {
		graph[v] = append(graph[v], i)
	}
	// 找寻最大的环
	pairs, result1 := countMaxLoop(favorite, graph)
	result2 := countTwoExtra(pairs, graph)
	return max(result1, result2)
}

func countTwoExtra(pairs [][]int, graph [][]int) int {
	res := 0
	n := len(graph)
	visited := make([]bool, n)
	var dfs func(i int) int
	dfs = func(i int) int {
		if visited[i] {
			return 0
		}
		res := 0
		visited[i] = true
		for _, v := range graph[i] {
			this := dfs(v)
			if this > res {
				res = this
			}
		}
		return res + 1
	}
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		visited[b] = true
		res += dfs(a)
		visited[b] = false
		visited[a] = true
		res += dfs(b)
		visited[a] = false
	}
	return res
}

func countMaxLoop(favorite []int, graph [][]int) (pairs [][]int, result int) {
	n := len(graph)
	// 0未遍历 1正在遍历中 2遍历结束
	state := make([]int, n)
	var dfs func(i int, count int)
	dfs = func(i int, count int) {
		if state[i] == 1 {
			result = max(result, count)
			if count == 2 {
				pairs = append(pairs, []int{i, favorite[i]})
			}
			return
		} else if state[i] == 0 {
			state[i] = 1
			for _, v := range graph[i] {
				dfs(v, count+1)
			}
			state[i] = 2
		}
	}
	for i := 0; i < n; i++ {
		dfs(i, 0)
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
