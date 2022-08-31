package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 18:53

 * @Note:	https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
			1372. Longest ZigZag Path in a Binary Tree
			1372. 二叉树中的最长交错路径
 **/

func longestZigZag(root *TreeNode) (result int) {
	var dfs func(root *TreeNode) (int, int)
	// int1 是先左拐 int2 是先右拐
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return -1, -1
		}
		left1, left2 := dfs(root.Left)
		right1, right2 := dfs(root.Right)
		result = max(result, left1, right2, 1+left2, 1+right1)
		return 1 + left2, 1 + right1
	}
	dfs(root)
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
