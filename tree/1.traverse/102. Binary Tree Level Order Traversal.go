package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 0:29

 * @Note:	https://leetcode.cn/problems/binary-tree-level-order-traversal/
			102. Binary Tree Level Order Traversal
			102. 二叉树的层序遍历
 **/

//老题了 在基础算法-BFS章节做过 很简单 就不接着写了
func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	pre := []*TreeNode{root}
	var next []*TreeNode
	for len(pre) != 0 {
		var this []int
		for _, node := range pre {
			this = append(this, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		result = append(result, this)
		pre = next
		next = []*TreeNode{}
	}
	return
}
