package main

import (
	"container/heap"
	"math"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 0:58

 * @Note:	https://leetcode.cn/problems/cheapest-flights-within-k-stops/
			787. Cheapest Flights Within K Stops
			787. K 站中转内最便宜的航班
 **/

// 迪杰斯特拉算法 未剪枝
//
//type PQ [][]int
//
//func (p *PQ) Len() int {
//	return len(*p)
//}
//
//func (p *PQ) Less(i, j int) bool {
//	return (*p)[i][2] < (*p)[j][2]
//}
//
//func (p *PQ) Swap(i, j int) {
//	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
//}
//
//func (p *PQ) Push(x interface{}) {
//	*p = append(*p, x.([]int))
//}
//
//func (p *PQ) Pop() interface{} {
//	n := p.Len()
//	out := (*p)[n-1]
//	*p = (*p)[:n-1]
//	return out
//}
//
//
//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	graph:=make(map[int][][2]int)
//	pq:=&PQ{}
//	for _, flight := range flights {
//		from,to,cost:=flight[0],flight[1],flight[2]
//		graph[from]=append(graph[from],[2]int{to,cost})
//		if from==src{
//			flight=append(flight,k)
//			heap.Push(pq,flight)
//		}
//	}
//	for pq.Len()!=0{
//		flight:= heap.Pop(pq).([]int)
//		_,to,cost,k:=flight[0],flight[1],flight[2],flight[3]
//		if dst==to{
//			return cost
//		}
//		if to!=src && k>0{
//			for _, e := range graph[to] {
//				heap.Push(pq,[]int{to,e[0],cost+e[1],k-1})
//			}
//		}
//	}
//	return -1
//}

// 抄代码是木有前途的！！！！！！！！！！！ 在用的时候会暴露的 ！！！！！！！（吐槽我之前）
//	Bellman-ford

//const MAX=100000
//func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
//	distances := make([]int, n)
//	for i := 0; i < n; i++ {
//		distances[i]=MAX
//	}
//	distances[src]=0
//	for i := 0; i <= k; i++ {
//		// 如果此处没有 copy的话 可能就会导致 多走
//		//for _, flight := range flights {
//		//	from,to,cost:=flight[0],flight[1],flight[2]
//		//	newCost:=distances[from]+cost
//		//	if newCost<distances[to]{
//		//		distances[to]=newCost
//		//	}
//		//}
//		next:=make([]int,n)
//		copy(next,distances)
//		for _, flight := range flights {
//			from,to,cost:=flight[0],flight[1],flight[2]
//			newCost:=distances[from]+cost
//			if newCost<next[to]{
//				next[to]=newCost
//			}
//		}
//		distances=next
//	}
//	if distances[dst]==MAX{
//		return -1
//	}else{
//		return distances[dst]
//	}
//}

// 迪杰斯特拉算法 剪枝版本
// 粗枝叶的剪枝

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

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make(map[int][][2]int)
	pq := &PQ{}
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]
		graph[from] = append(graph[from], [2]int{to, cost})
		if from == src {
			flight = append(flight, k)
			heap.Push(pq, flight)
		}
	}
	visited := make([]int, n)
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = math.MaxInt64
		distance[i] = math.MaxInt64
	}
	distance[src] = 0
	visited[src] = 0
	for pq.Len() != 0 {
		flight := heap.Pop(pq).([]int)
		_, to, cost, k := flight[0], flight[1], flight[2], flight[3]
		if dst == to {
			return cost
		}
		if k > 0 {
			for _, e := range graph[to] {
				newCost := cost + e[1]
				visit := k - 1
				if newCost >= distance[e[0]] && visited[e[0]] >= visit {
					continue
				}
				distance[e[0]] = min(distance[e[0]], newCost)
				visited[e[0]] = min(visited[e[0]], visit)
				heap.Push(pq, []int{to, e[0], cost + e[1], visit})
			}
		}
	}
	return -1
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
