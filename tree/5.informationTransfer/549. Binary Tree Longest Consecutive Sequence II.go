package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 19:30

 * @Note:	https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence-ii/
			549. Binary Tree Longest Consecutive Sequence II
			549. 二叉树中最长的连续序列
 **/

// 单纯递增或者递减
func longestConsecutive(root *TreeNode) (result int) {
	// 第一个值是递减 第二个值是递增
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		desc, asc := 1, 1
		var leftDesc, leftAsc, rightAsc, rightDesc int
		if root.Left != nil {
			leftDesc, leftAsc = dfs(root.Left)
			if root.Val-1 == root.Left.Val {
				desc = leftDesc + 1
			}
			if root.Val+1 == root.Left.Val {
				asc = leftAsc + 1
			}
		}
		if root.Right != nil {
			rightDesc, rightAsc = dfs(root.Right)
			if root.Val-1 == root.Right.Val {
				desc = max(desc, rightDesc+1)
				if root.Left != nil && root.Left.Val == root.Val+1 {
					result = max(result, leftAsc+rightDesc+1)
				}
			}
			if root.Val+1 == root.Right.Val {
				asc = max(asc, rightAsc+1)
				if root.Left != nil && root.Left.Val == root.Val-1 {
					result = max(result, leftDesc+rightAsc+1)
				}
			}
		}
		result = max(result, desc, asc)
		return desc, asc
	}
	dfs(root)
	return
}

func max(ints ...int) int {
	result := ints[0]
	for _, v := range ints[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

//func longestConsecutive(root *TreeNode) (result int) {
//	// 第一个值是递减 第二个值是递增
//	var dfs func(root *TreeNode) int
//	dfs = func(root *TreeNode) int {
//		var left,right int
//		this:=1
//		returnNum:=1
//		if root.Left!=nil{
//			left = dfs(root.Left)
//			if root.Left.Val==root.Val-1 || root.Left.Val==root.Val+1{
//				this+=left
//				returnNum+=left
//			}
//		}
//		if root.Right!=nil{
//			right= dfs(root.Right)
//			if root.Right.Val==root.Val-1 || root.Right.Val==root.Val+1{
//				this+=right
//				returnNum=max(returnNum,1+right)
//			}
//		}
//		result=max(result,this,left,right)
//		return returnNum
//	}
//	dfs(root)
//	return
//}
