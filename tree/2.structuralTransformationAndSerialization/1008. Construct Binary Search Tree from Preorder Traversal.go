package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:31

 * @Note:	https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal/
			1008. Construct Binary Search Tree from Preorder Traversal
			1008. 前序遍历构造二叉搜索树
 **/

func bstFromPreorder(preorder []int) *TreeNode {
	n := len(preorder)
	i := 0
	var dfs func(min, max int) *TreeNode
	dfs = func(min, max int) *TreeNode {
		if i >= n {
			return nil
		}
		if min < preorder[i] && max > preorder[i] {
			root := &TreeNode{Val: preorder[i]}
			i++
			root.Left = dfs(min, root.Val)
			root.Right = dfs(root.Val, max)
			return root
		}
		return nil
	}
	return dfs(math.MinInt64, math.MaxInt)
}
