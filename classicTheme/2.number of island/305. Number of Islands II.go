package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 16:49

 * @Note:	https://leetcode.cn/problems/number-of-islands-ii/
			305. Number of Islands II
			305. 岛屿数量 II
 **/
// 离线 UF

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

func numIslands2(m int, n int, positions [][]int) []int {
	num := 0
	length := m * n
	uf := NewUF(length)
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}
	result := make([]int, len(positions))
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for i, position := range positions {
		x, y := position[0], position[1]
		if table[x][y] == 0 {
			table[x][y] = 1
			num++
			for _, dir := range dirs {
				newX, newY := x+dir[0], y+dir[1]
				if newX < 0 || newY < 0 || newX >= m || newY >= n {
					continue
				}
				if table[newX][newY] == 1 {
					if !uf.checkBinding(newX*n+newY, x*n+y) {
						uf.bind(newX*n+newY, x*n+y)
						num--
					}
				}
			}
		}
		result[i] = num
	}
	return result
}
