package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 13:00

 * @Note:	https://leetcode.cn/problems/deepest-leaves-sum/
			1302. Deepest Leaves Sum
			1302. 层数最深叶子节点的和
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层次遍历

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := []*TreeNode{root}
	next := make([]*TreeNode, 0)
	for {
		sum := 0
		for _, node := range list {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			sum += node.Val
		}
		if len(next) == 0 {
			return sum
		} else {
			list = next
			next = make([]*TreeNode, 0)
		}
	}
}
