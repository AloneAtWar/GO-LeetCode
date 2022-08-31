package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 18:23

 * @Note:	https://leetcode.cn/problems/binary-tree-maximum-path-sum/
			124. Binary Tree Maximum Path Sum
			124. 二叉树中的最大路径和
 **/

func maxPathSum(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	// 第一个结果是一定要有root节点 且没有在 root节点处拐弯
	// 第二个结果是 可以没有root节点 但考虑过拐弯的情况
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, math.MinInt64
		}
		left1, left2 := dfs(root.Left)
		right1, right2 := dfs(root.Right)
		return max(left1, right1, 0) + root.Val, max(max(left1, 0)+root.Val+max(right1, 0), left2, right2)
	}
	_, result := dfs(root)
	return result
}

func max(ints ...int) int {
	result := ints[0]
	for _, v := range ints[1:] {
		if v > result {
			result = v
		}
	}
	return result
}
