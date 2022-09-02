package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 12:13

 * @Note:

 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//标准的信息传递类题目

func longestUnivaluePath(root *TreeNode) (result int) {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		curr := 1
		left1, left2 := dfs(root.Left)
		right1, right2 := dfs(root.Right)
		if left1 != root.Val {
			left2 = 0
		}
		if right1 != root.Val {
			right2 = 0
		}
		this := left2 + right2
		if result < this {
			result = this
		}
		return root.Val, curr + max(left2, right2)
	}
	dfs(root)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
