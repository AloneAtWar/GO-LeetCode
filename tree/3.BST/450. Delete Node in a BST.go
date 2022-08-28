package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 2:18

 * @Note:	https://leetcode.cn/problems/delete-node-in-a-bst/
			450. Delete Node in a BST
			450. 删除二叉搜索树中的节点
 **/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return deleteRoot(root)
	}
	curr := root
	for (curr.Val > key && curr.Left != nil) || (curr.Val < key && curr.Right != nil) {
		if curr.Val > key && curr.Left != nil {
			if curr.Left.Val == key {
				curr.Left = deleteRoot(curr.Left)
				break
			} else {
				curr = curr.Left
			}
		} else {
			if curr.Right.Val == key {
				curr.Right = deleteRoot(curr.Right)
				break
			} else {
				curr = curr.Right
			}
		}
	}
	return root
}

func deleteRoot(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}
	curr := node.Right
	for curr.Left != nil {
		curr = curr.Left
	}
	curr.Left = node.Left
	return node.Right
}
