package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 3:01

 * @Note:	https://leetcode.cn/problems/validate-binary-search-tree/
			98. Validate Binary Search Tree
			98. 验证二叉搜索树
 **/
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
