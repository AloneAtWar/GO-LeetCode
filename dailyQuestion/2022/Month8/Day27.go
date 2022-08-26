package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/27 2:18

 * @Note:	https://leetcode.cn/problems/maximum-width-of-binary-tree/
			662. Maximum Width of Binary Tree
			662. 二叉树最大宽度
 **/

func widthOfBinaryTree(root *TreeNode) (result int) {
	if root == nil {
		return 0
	}
	maxHeight := 0
	max := map[int]int{}
	min := map[int]int{}
	var dfs func(root *TreeNode, height, local int)
	dfs = func(root *TreeNode, height, local int) {
		if root == nil {
			return
		}
		if max[height] == 0 {
			max[height] = local
			min[height] = local
			maxHeight++
		} else if max[height] < local {
			max[height] = local
		} else if min[height] > local {
			min[height] = local
		}
		dfs(root.Left, height+1, local*2)
		dfs(root.Right, height, local*2+1)
	}
	dfs(root, 0, 1)
	for i := 1; i <= maxHeight; i++ {
		this := max[root.Val] - min[root.Val]
		if this > result {
			result = this
		}
	}
	return
}
