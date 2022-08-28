package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:49

 * @Note:	https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
			426. 将二叉搜索树转化为排序的双向链表
			426. Convert Binary Search Tree to Sorted Doubly Linked List
 **/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var result, pre *Node = nil, nil
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root.Left != nil {
			dfs(root.Left)
		}
		if pre != nil {
			pre.Right = root
			root.Left = pre
		}
		pre = root
		if result == nil {
			result = root
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	pre.Right = result
	result.Left = pre
	return result
}
