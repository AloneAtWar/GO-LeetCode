package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 0:55

 * @Note:	https://leetcode.cn/problems/subsets-ii/
			90. Subsets II
			90. 子集 II
 **/

func subsetsWithDup(nums []int) (result [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var dfs func(i int)
	var curr []int
	dfs = func(i int) {
		result = append(result, append([]int(nil), curr...))
		n := len(curr)
		for j := i; j < len(nums); j++ {
			// j==i 说明是 为第一个考虑的元素
			// 同时 j!=i 保证了 j-1 不可能<0 因为 j==0时
			// i也一定为0   j>=i i>=0
			if j == i || nums[j] != nums[j-1] {
				curr = append(curr, nums[j])
				dfs(j + 1)
				curr = curr[:n]
			}
		}
	}
	dfs(0)
	return
}
