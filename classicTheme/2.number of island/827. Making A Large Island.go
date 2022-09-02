package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 18:36

 * @Note:	https://leetcode.cn/problems/making-a-large-island/
			827. Making A Large Island
			827. 最大人工岛
 **/

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

func (u *UF) findSize(i int) int {
	return u.size[u.findFather(i)]
}
func (u *UF) checkBinding(i, j int) bool {
	return u.findFather(i) == u.findFather(j)
}

func largestIsland(grid [][]int) (result int) {
	m, n := len(grid), len(grid[0])
	uf := NewUF(m * n)
	var dfs func(i, j int)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	dfs = func(i, j int) {
		grid[i][j] = -1
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			if newI < 0 || newJ < 0 || newI >= m || newJ >= n || grid[newI][newJ] != 1 {
				continue
			}
			uf.bind(i*n+j, newI*n+newJ)
			dfs(newI, newJ)
		}
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				size := uf.findSize(i*n + j)
				if size > result {
					result = size
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				thisSize := 1
				has := make(map[int]int)
				for _, dir := range dirs {
					newI, newJ := i+dir[0], j+dir[1]
					if newI < 0 || newJ < 0 || newI >= m || newJ >= n || grid[newI][newJ] != -1 {
						continue
					}
					father := uf.findFather(newI*n + newJ)
					has[father] = uf.findSize(father)
				}
				for _, v := range has {
					thisSize += v
				}
				if thisSize > result {
					result = thisSize
				}
			}
		}

	}
	return
}
