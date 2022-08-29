package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 22:42

 * @Note:	https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iii/
			1650. Lowest Common Ancestor of a Binary Tree III
			1650. 二叉树的最近公共祖先 III
 **/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	a, b := p, q
	for p != q {
		if p == nil {
			p = a
		} else {
			p = p.Parent
		}
		if q == nil {
			q = b
		} else {
			q = q.Parent
		}
	}
	return p
}
