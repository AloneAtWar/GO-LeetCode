package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 21:18

 * @Note:	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
			236. Lowest Common Ancestor of a Binary Tree
			236. 二叉树的最近公共祖先
 **/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left == nil {
		return right
	} else {
		return left
	}
}
