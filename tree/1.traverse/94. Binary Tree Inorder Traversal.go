package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:58

 * @Note:	https://leetcode.cn/problems/binary-tree-inorder-traversal/
			94. Binary Tree Inorder Traversal
			94. 二叉树的中序遍历
 **/

// 依旧是非递归
func inorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	curr := root
	var stack = []*TreeNode{}
	for curr != nil || len(stack) != 0 {
		if curr == nil {
			curr = stack[len(stack)-1]
			result = append(result, curr.Val)
			curr = curr.Right
			stack = stack[:len(stack)-1]
		} else {
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.Left
			}
		}
	}
	return
}
