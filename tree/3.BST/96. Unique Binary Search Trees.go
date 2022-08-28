package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 3:39

 * @Note:	https://leetcode.cn/problems/unique-binary-search-trees/
			96. Unique Binary Search Trees
			96. 不同的二叉搜索树
 **/

func numTrees(n int) int {
	dp := map[int]int{}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 1
		}
		if dp[i] > 0 {
			return dp[i]
		}
		result := 0
		for j := 1; j <= i; j++ {
			result += dfs(j-1) * dfs(i-j)
		}
		dp[i] = result
		return result
	}
	return dfs(n)
}
