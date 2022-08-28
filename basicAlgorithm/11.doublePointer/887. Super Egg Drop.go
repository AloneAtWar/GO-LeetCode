package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/28 22:07

 * @Note:	https://leetcode.cn/problems/super-egg-drop/
			887. Super Egg Drop
			887. 鸡蛋掉落
 **/

func superEggDrop(k int, n int) int {
	table := make(map[int][]int)
	for i := 0; i <= k; i++ {
		table[i] = make([]int, n+1)
	}
	var dfs func(k, n int) int
	dfs = func(k, n int) int {
		if n <= 1 || k == 1 {
			return n
		}
		if table[k][n] > 0 {
			return table[k][n]
		}
		low, hight := 1, n
		res := n
		for low <= hight {
			mid := (low + hight) >> 1
			left := dfs(k-1, mid-1)
			right := dfs(k, n-mid)
			res = min(res, max(left, right)+1)
			if left == right {
				break
			} else if left > right {
				hight = mid - 1
			} else {
				low = mid + 1
			}
		}
		table[k][n] = res
		return res
	}
	return dfs(k, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
