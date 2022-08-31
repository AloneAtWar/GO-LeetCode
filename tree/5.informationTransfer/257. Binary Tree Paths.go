package main

import (
	"strconv"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 17:58

 * @Note:	https://leetcode.cn/problems/binary-tree-paths/
			257. Binary Tree Paths
			257. 二叉树的所有路径
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) (result []string) {
	if root == nil {
		return
	}
	curr := strconv.Itoa(root.Val)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left == nil && root.Right == nil {
			result = append(result, curr)
			return
		}
		if root.Left != nil {
			next := "->" + strconv.Itoa(root.Left.Val)
			temp := curr
			curr += next
			dfs(root.Left)
			curr = temp
		}
		if root.Right != nil {
			next := "->" + strconv.Itoa(root.Right.Val)
			temp := curr
			curr += next
			dfs(root.Right)
			curr = temp
		}
	}
	dfs(root)
	return
}
