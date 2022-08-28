package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:38

 * @Note:	https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
			105. 从前序与中序遍历序列构造二叉树
			105. Construct Binary Tree from Preorder and Inorder Traversal
 **/

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	return build(preorder, 0, n-1, inorder, 0, n-1)
}

func build(preorder []int, left1, right1 int, inOrder []int, left2, right2 int) *TreeNode {
	if left1 > right1 || left2 > right2 {
		return nil
	}
	val := preorder[left1]
	root := &TreeNode{Val: val}
	if left1 == right1 {
		return root
	}
	for i := left2; i <= right2; i++ {
		if val == inOrder[i] {
			root.Left = build(preorder, left1+1, left1+i-left2, inOrder, left2, i-1)
			root.Right = build(preorder, left1+i-left2+1, right1, inOrder, i+1, right2)
		}
	}
	return root
}
