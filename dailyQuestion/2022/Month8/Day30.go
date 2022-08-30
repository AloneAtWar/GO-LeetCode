package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 1:45

 * @Note:	https://leetcode.cn/problems/maximum-binary-tree-ii/
			998. Maximum Binary Tree II
			998. 最大二叉树 II
 **/

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		newNode := &TreeNode{Val: val}
		newNode.Left = root
		return newNode
	} else {
		root.Right = insertIntoMaxTree(root.Right, val)
		return root
	}
}
