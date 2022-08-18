package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 2:40

 * @Note:	https://leetcode.cn/problems/constrained-subsequence-sum/
			1425. Constrained Subsequence Sum
			1425. 带限制的子序列和
 **/

//这题的思路真的非常的绝
// 从子序列的地方 着手 考虑到 以某元素结尾
// 由此想到分治 分治想到 记忆化 想到dp
// 之后想到用 单调队列
func constrainedSubsetSum(nums []int, k int) int {
	result := math.MinInt64
	n := len(nums)
	var list []int
	sum := make([]int, n)
	for i, num := range nums {
		sum[i] = num
		if len(list) != 0 {
			sum[i] += sum[list[0]]
		}
		if sum[i] > result {
			result = sum[i]
		}
		if len(list) != 0 && i-list[0] >= k {
			list = list[1:]
		}
		for len(list) != 0 && sum[list[len(list)-1]] <= sum[i] {
			list = list[:len(list)-1]
		}
		if sum[i] > 0 {
			list = append(list, i)
		}
	}
	return result
}
