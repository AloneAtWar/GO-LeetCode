package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/3 3:18

 * @Note:	https://leetcode.cn/problems/maximum-length-of-pair-chain/
			646. Maximum Length of Pair Chain
			646. 最长数对链
 **/

// 一直没有读懂题目 把题目相乘扫描线 区间类题目 轻松贪心过
//func findLongestChain(pairs [][]int) int {
//	sort.Slice(pairs, func(i, j int) bool {
//		return pairs[i][0]<pairs[j][0]
//	})
//	res,pre:=1,pairs[0][1]
//	for _, pair := range pairs[1:] {
//		if pair[0]>pre{
//			res++
//			pre=pair[1]
//		}else{
//			pre=min(pre,pair[1])
//		}
//	}
//	return res
//}
//
//func min(a,b int)int{
//	if a>b{
//		return b
//	}
//	return a
//}

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })
	n := len(pairs)
	dp := make([]int, n)
	for i, p := range pairs {
		dp[i] = 1
		for j, q := range pairs[:i] {
			if p[0] > q[1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
