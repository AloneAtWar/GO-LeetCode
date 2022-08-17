package main

import (
	"container/heap"
	"sort"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 16:44

 * @Note:	https://leetcode.cn/problems/the-skyline-problem/
			218. The Skyline Problem
			218. 天际线问题
 **/
type Point struct {
	x int
	h int
	t int // 0 代表开始 1代表结束
}

type height []int

func (m *height) Len() int {
	return len(*m)
}

func (m *height) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *height) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *height) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *height) Pop() interface{} {
	out := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return out
}

func getSkyline(buildings [][]int) (result [][]int) {
	var points []*Point
	for _, building := range buildings {
		points = append(points, &Point{building[0], building[2], 0})
		points = append(points, &Point{building[1], building[2], 1})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		if points[i].t != points[j].t {
			return points[i].t == 0
		}
		if points[i].t == 0 {
			return points[i].h > points[j].h
		} else {
			return points[i].h < points[j].h
		}
		// 如果两个都是开始的话 那么应该 高的先
		// 如果两个都是结束的话 那么应该 矮的先
		// 如果一个开始一个结束的话 那么应该 开始先 （ 开始需要接住 结束 让他不要直接掉下去
	})
	//现在的高度
	curr := 0
	// 用来看高度是否合法
	legal := map[int]int{}
	h := &height{}
	for _, point := range points {
		if point.t == 0 {
			if point.h > curr {
				curr = point.h
				result = append(result, []int{point.x, curr})
			}
			heap.Push(h, point.h)
			legal[point.h]++
		} else {
			legal[point.h]--
			he := 0
			for h.Len() != 0 {
				if legal[(*h)[0]] > 0 {
					he = (*h)[0]
				} else {
					heap.Pop(h)
				}
			}
			if curr > he {
				result = append(result, []int{point.x, he})
			}
			curr = he
		}
	}
	return
}
