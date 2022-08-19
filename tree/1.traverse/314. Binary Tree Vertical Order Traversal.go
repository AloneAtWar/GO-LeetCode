package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 1:42

 * @Note:	https://leetcode.cn/problems/binary-tree-vertical-order-traversal/
			314. Binary Tree Vertical Order Traversal
			314. 二叉树的垂直遍历
 **/

// 尝试了半天的 dfs 之后再排序的方法
// 发现根本不行
// 老老实实使用 bfs 吧
func verticalOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	table := make(map[int][]int)
	min := 0
	max := 0
	type node struct {
		root      *TreeNode
		localhost int
	}
	list := []node{{root, 0}}
	for len(list) != 0 {
		out := list[0]
		list = list[1:]
		table[out.localhost] = append(table[out.localhost], out.root.Val)
		if out.root.Left != nil {
			list = append(list, node{out.root.Left, out.localhost - 1})
			if out.localhost-1 < min {
				min = out.localhost - 1
			}
		}
		if out.root.Right != nil {
			list = append(list, node{out.root.Right, out.localhost + 1})
			if out.localhost+1 > max {
				max = out.localhost + 1
			}
		}
	}
	for i := min; i <= max; i++ {
		result = append(result, table[i])
	}
	return
}
