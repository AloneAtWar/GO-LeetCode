package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 14:43

 * @Note:	https://leetcode.cn/problems/critical-connections-in-a-network/
			1192. Critical Connections in a Network
			1192. 查找集群内的「关键连接」
 **/

func criticalConnections(n int, connections [][]int) (result [][]int) {
	graph := make([][]int, n)
	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], connection[1])
		graph[connection[1]] = append(graph[connection[1]], connection[0])
	}
	var dfs func(i int, father int)
	j := 1
	low := make([]int, n)
	dfn := make([]int, n)
	dfs = func(i int, father int) {
		low[i] = j
		dfn[i] = j
		j++
		for _, v := range graph[i] {
			if v != father {
				if low[v] == 0 {
					dfs(v, i)
				}
				if dfn[v] < dfn[i] {
					dfn[i] = dfn[v]
				}
			}
		}
		if dfn[i] == low[i] && father != -1 {
			result = append(result, []int{father, i})
		}
	}
	dfs(0, -1)
	return
}
