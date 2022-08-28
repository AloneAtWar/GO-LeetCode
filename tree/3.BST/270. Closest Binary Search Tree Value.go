package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 2:07

 * @Note:	https://leetcode.cn/problems/closest-binary-search-tree-value/
			270. Closest Binary Search Tree Value
			270. 最接近的二叉搜索树值
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	this := float64(root.Val)
	if this > target {
		if root.Left != nil {
			value := closestValue(root.Left, target)
			if abs(float64(value)-target) < this-target {
				return value
			} else {
				return root.Val
			}
		} else {
			return root.Val
		}
	} else {
		if root.Right != nil {
			value := closestValue(root.Right, target)
			if abs(float64(value)-target) < target-this {
				return value
			} else {
				return root.Val
			}
		} else {
			return root.Val
		}
	}
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}
