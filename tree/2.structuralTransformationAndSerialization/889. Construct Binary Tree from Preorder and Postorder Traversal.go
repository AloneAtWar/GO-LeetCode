package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 1:02

 * @Note:	https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
			889. Construct Binary Tree from Preorder and Postorder Traversal
			889. 根据前序和后序遍历构造二叉树
 **/

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	var build func(left1, right1, left2, right2 int) *TreeNode
	build = func(left1, right1, left2, right2 int) *TreeNode {
		if left1 > right1 {
			return nil
		}
		val := preorder[left1]
		root := &TreeNode{Val: val}
		if left1 == right1 {
			return root
		}
		find := left1 + 1
		for i := left2; i <= right2-1; i++ {
			if postorder[i] == preorder[find] {
				root.Left = build(left1+1, left1+1+i-left2, left2, i)
				root.Right = build(left1+1+i-left2+1, right1, i+1, right2-1)
			}
		}
		return root
	}
	return build(0, n-1, 0, n-1)
}
