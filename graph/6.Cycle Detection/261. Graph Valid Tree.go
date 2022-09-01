package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 18:03

 * @Note:	https://leetcode.cn/problems/graph-valid-tree/
			261. Graph Valid Tree
			261. 以图判树
 **/

func validTree(n int, edges [][]int) bool {
	valid := true
	graph := buildGraph(n, edges)
	visited := make(map[int]bool)
	var dfs func(i int, father int)
	dfs = func(i int, father int) {
		if !valid {
			return
		}
		visited[i] = true
		for _, v := range graph[i] {
			if v != father {
				if visited[v] {
					valid = false
					return
				} else {
					dfs(v, i)
				}
			}
		}
	}
	dfs(0, -1)
	return valid && len(visited) == n
}

func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}
