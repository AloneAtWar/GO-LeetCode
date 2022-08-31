package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 22:57

 * @Note:	https://leetcode.cn/problems/connecting-cities-with-minimum-cost/
			1135. Connecting Cities With Minimum Cost
			1135. 最低成本联通所有城市
 **/

//这里就使用UF+Kruskal
func minimumCost(n int, connections [][]int) int {
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})
	result := 0
	uf := NewUF(n)
	for _, connection := range connections {
		from, to, cost := connection[0]-1, connection[1]-1, connection[2]
		if !uf.checkBinding(from, to) {
			result += cost
			uf.bind(from, to)
			if uf.num == 1 {
				break
			}
		}
	}
	if uf.num != 1 {
		return -1
	} else {
		return result
	}
}

type UF struct {
	father []int
	size   []int
	num    int
}

func NewUF(n int) *UF {
	father := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		father[i] = i
		size[i] = 1
	}
	return &UF{father: father, size: size, num: n}
}

func (u *UF) findFather(i int) int {
	for u.father[i] != i {
		i = u.father[i]
	}
	return i
}

func (u *UF) bind(i, j int) {
	iFahter, jFather := u.findFather(i), u.findFather(j)
	if iFahter != jFather {
		if u.size[iFahter] > u.size[jFather] {
			u.father[jFather] = iFahter
			u.size[iFahter] += u.size[jFather]
		} else {
			u.father[iFahter] = jFather
			u.size[jFather] += u.size[iFahter]
		}
		u.num--
	}
}

func (u *UF) checkBinding(i, j int) bool {
	return u.findFather(i) == u.findFather(j)
}
