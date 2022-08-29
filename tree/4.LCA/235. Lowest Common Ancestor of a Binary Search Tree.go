package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 21:12

 * @Note:	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
			235. Lowest Common Ancestor of a Binary Search Tree
			235. 二叉搜索树的最近公共祖先
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	for {
		if root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val {
			root = root.Right
		} else {
			return root
		}
	}
}
