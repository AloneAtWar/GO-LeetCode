package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 3:04

 * @Note:	https://leetcode.cn/problems/binary-search-tree-iterator/
			173. Binary Search Tree Iterator
			173. 二叉搜索树迭代器
 **/

type BSTIterator struct {
	stack []*TreeNode
	curr  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{curr: root}
}

func (this *BSTIterator) Next() int {
	if this.curr != nil {
		for this.curr != nil {
			this.stack = append(this.stack, this.curr)
			this.curr = this.curr.Left
		}
	}
	out := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.curr = out.Right
	return out.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.curr != nil
}
