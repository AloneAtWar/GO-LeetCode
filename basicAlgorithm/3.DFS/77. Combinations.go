package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 1:14

 * @Note:	https://leetcode.cn/problems/combinations/
			77. Combinations
			77. 组合
 **/

func combine(n int, k int) (result [][]int) {
	var dfs func(i int)
	var curr []int
	dfs = func(i int) {
		size := len(curr)
		if size == k {
			result = append(result, append([]int(nil), curr...))
			return
		}
		for j := i; j <= n; j++ {
			curr = append(curr, j)
			dfs(j + 1)
			curr = curr[:size]
		}
	}
	dfs(1)
	return
}
