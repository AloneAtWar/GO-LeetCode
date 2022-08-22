package main

import (
	"fmt"
	"math"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/22 15:12

 * @Note:

 **/

func printTree(root *TreeNode) [][]string {
	var dfs func(root *TreeNode, n int)
	height := 0
	dfs = func(root *TreeNode, n int) {
		if root == nil {
			return
		}
		if n > height {
			height = n
		}
		dfs(root.Left, n+1)
		dfs(root.Right, n+1)
	}
	dfs(root, 0)
	result := make([][]string, height+1)
	m, n := height+1, pow(2, height+1)-1
	for i := 0; i < m; i++ {
		result[i] = make([]string, n)
	}
	type node struct {
		root *TreeNode
		i, j int
	}
	nodes := []*node{{root, 0, (n - 1) / 2}}
	for len(nodes) != 0 {
		out := nodes[0]
		nodes = nodes[1:]
		result[out.i][out.j] = fmt.Sprintf("%d", out.root.Val)
		if out.root.Left != nil {
			nodes = append(nodes, &node{out.root.Left, out.i + 1, out.j - pow(2, height-out.i-1)})
		}
		if out.root.Right != nil {
			nodes = append(nodes, &node{out.root.Right, out.i + 1, out.j + pow(2, height-out.i-1)})
		}
	}
	return result
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
