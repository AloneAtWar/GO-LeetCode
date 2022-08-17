package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 19:39

 * @Note:	https://leetcode.cn/problems/minimum-depth-of-binary-tree/
			111. Minimum Depth of Binary Tree
			111. 二叉树的最小深度
 **/

func minDepth(root *TreeNode) (result int) {
	if root == nil {
		return 0
	}
	pre := []*TreeNode{root}
	var next []*TreeNode
	for len(pre) != 0 {
		result++
		for _, node := range pre {
			if node.Left == nil && node.Right == nil {
				return
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		pre = next
		next = []*TreeNode{}
	}
	return
}
