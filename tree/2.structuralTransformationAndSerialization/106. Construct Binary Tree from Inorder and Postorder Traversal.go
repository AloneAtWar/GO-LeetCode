package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:52

 * @Note:	https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
			106. Construct Binary Tree from Inorder and Postorder Traversal
			106. 从中序与后序遍历序列构造二叉树
 **/

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	var build func(left1, right1, left2, right2 int) *TreeNode
	build = func(left1, right1, left2, right2 int) *TreeNode {
		if left1 > right1 {
			return nil
		}
		val := postorder[right2]
		result := &TreeNode{Val: val}
		if left1 == right1 {
			return result
		}
		for i := left1; i <= right1; i++ {
			if inorder[i] == val {
				result.Left = build(left1, i-1, left2, left2+i-left1-1)
				result.Right = build(i+1, right1, left2+i-left1, right2-1)
				break
			}
		}
		return result
	}
	return build(0, n-1, 0, n-1)
}
