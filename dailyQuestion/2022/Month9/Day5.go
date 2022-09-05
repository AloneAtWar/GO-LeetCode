package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/5 11:40

 * @Note:	https://leetcode.cn/problems/find-duplicate-subtrees/
			652. Find Duplicate Subtrees
			652. 寻找重复的子树
 **/

func findDuplicateSubtrees(root *TreeNode) (ans []*TreeNode) {
	table := make(map[string]int)
	var helper func(root *TreeNode) string
	helper = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		result := fmt.Sprintf("%s-%d-%s", helper(root.Left), root.Val, helper(root.Right))
		if table[result] == 1 {
			ans = append(ans, root)
		}
		table[result]++
		return result
	}
	helper(root)
	return
}
