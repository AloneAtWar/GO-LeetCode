package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 23:14

 * @Note:	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iv/
			1676. Lowest Common Ancestor of a Binary Tree IV
			1676. 二叉树的最近公共祖先 IV
 **/

// 因为题目不支持GO语言作答所以我就换一个语言来进行测试
func lowestCommonAncestor1676(root *TreeNode, nodes []*TreeNode) *TreeNode {
	table := make(map[int]bool)
	for _, node := range nodes {
		table[node.Val] = true
	}
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if table[root.Val] {
			return root
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left != nil && right != nil {
			return root
		}
		if left == nil {
			return right
		} else {
			return left
		}
	}
	return dfs(root)
}
