package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 18:18

 * @Note:	https://leetcode.cn/problems/count-good-nodes-in-binary-tree/
			1448. Count Good Nodes in Binary Tree
			1448. 统计二叉树中好节点的数目
 **/

func goodNodes(root *TreeNode) (result int) {
	if root == nil {
		return
	}
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			result++
			max = root.Val
		}
		dfs(root.Right, max)
		dfs(root.Left, max)
	}
	dfs(root, root.Val)
	return
}
