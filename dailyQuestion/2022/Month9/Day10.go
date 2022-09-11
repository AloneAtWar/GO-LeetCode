package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/11 12:23

 * @Note:	https://leetcode.cn/problems/trim-a-binary-search-tree/submissions/
			669. Trim a Binary Search Tree
			669. 修剪二叉搜索树
 **/

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val > high {
			return dfs(root.Left)
		} else if root.Val < low {
			return dfs(root.Right)
		} else {

			root.Left = dfs(root.Left)
			root.Right = dfs(root.Right)
			return root
		}
	}
	return dfs(root)
}
