package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 1:28

 * @Note:	https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
			107. Binary Tree Level Order Traversal II
			107. 二叉树的层序遍历 II
 **/

// 102 的代码 修修补补又是一年
func levelOrderBottom(root *TreeNode) [][]int {
	return reverse(levelOrder(root))
}

func reverse(result [][]int) [][]int {
	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}
