package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 3:12

 * @Note:	https://leetcode.cn/problems/recover-binary-search-tree/
			99. Recover Binary Search Tree.go
			99. 恢复二叉搜索树
 **/

func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
