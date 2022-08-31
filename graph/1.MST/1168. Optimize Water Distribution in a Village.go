package main

import "container/heap"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 23:48

 * @Note:	https://leetcode.cn/problems/optimize-water-distribution-in-a-village/
			1168. Optimize Water Distribution in a Village
			1168. 水资源分配优化
 **/

// 纯属炫技!  这道题用这个贼难写

// PQ+Prim

type PQ [][]int

func (p *PQ) Len() int {
	return len(*p)
}

func (p *PQ) Less(i, j int) bool {
	return (*p)[i][2] < (*p)[j][2]
}

func (p *PQ) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PQ) Push(x interface{}) {
	*p = append(*p, x.([]int))
}

func (p *PQ) Pop() interface{} {
	n := p.Len()
	out := (*p)[n-1]
	*p = (*p)[:n-1]
	return out
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) (result int) {
	pq := &PQ{}
	for i, well := range wells {
		*pq = append(*pq, []int{0, i + 1, well})
	}
	graph := make([][][2]int, n+1)
	for _, out := range pipes {
		from, to, cost := out[0], out[1], out[2]
		graph[from] = append(graph[from], [2]int{to, cost})
		graph[to] = append(graph[to], [2]int{from, cost})
	}
	heap.Init(pq)
	visited := make([]bool, n+1)
	for pq.Len() != 0 {
		out := heap.Pop(pq).([]int)
		_, to, cost := out[0], out[1], out[2]
		if visited[to] == true {
			continue
		}
		result += cost
		visited[to] = true
		for _, ints := range graph[to] {
			if !visited[ints[0]] {
				heap.Push(pq, []int{to, ints[0], ints[1]})
			}
		}
	}
	return
}
