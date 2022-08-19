package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:52

 * @Note:	https://leetcode.cn/problems/binary-tree-preorder-traversal/
			144. Binary Tree Preorder Traversal
			144. 二叉树的前序遍历
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 老题 基础算法-DFS 章节用递归的方式实现过一次 这次换 非递归来弄
// 这个真的是最简单的
func preorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return
}
