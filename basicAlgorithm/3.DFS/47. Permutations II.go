package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 1:10

 * @Note:	https://leetcode.cn/problems/permutations-ii/
			47. Permutations II
			47. 全排列 II
 **/

func permuteUnique(nums []int) (result [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var dfs func()
	table := map[int]bool{}
	n := len(nums)
	curr := make([]int, 0, len(nums))
	// 上一个安排的
	dfs = func() {
		pre := -99
		if len(curr) == n {
			result = append(result, append([]int(nil), curr...))
			return
		}
		size := len(curr)
		for i, num := range nums {
			if !table[i] && pre != num {
				curr = append(curr, num)
				table[i] = true
				dfs()
				table[i] = false
				curr = curr[:size]
				pre = i
			}
		}
	}
	dfs()
	return
}
