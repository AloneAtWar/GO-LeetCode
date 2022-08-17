package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 2:09

 * @Note:	https://leetcode.cn/problems/binary-tree-preorder-traversal/
			144. Binary Tree Preorder Traversal
			144. 二叉树的前序遍历
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (result []int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			result = append(result, root.Val)
			dfs(root.Left)

			dfs(root.Right)
		}
	}
	dfs(root)
	return
}
