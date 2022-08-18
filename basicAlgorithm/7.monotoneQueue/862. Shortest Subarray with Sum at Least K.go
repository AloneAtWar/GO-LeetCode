package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 2:26

 * @Note:	https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
			862. Shortest Subarray with Sum at Least K
			862. 和至少为 K 的最短子数组
 **/

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	var list []int
	result := math.MaxInt64
	// 2 -1 2
	// 2 1 3
	// 前缀合
	sum := make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	for i := 0; i <= n; i++ {
		for len(list) != 0 && sum[i]-sum[list[0]] >= k {
			currSize := i - list[0]
			if currSize < result {
				result = currSize
			}
			list = list[1:]
		}
		for len(list) != 0 && sum[i] <= sum[list[len(list)-1]] {
			list = list[:len(list)-1]
		}
		list = append(list, i)
	}
	if result == math.MaxInt64 {
		return -1
	} else {
		return result
	}
}
