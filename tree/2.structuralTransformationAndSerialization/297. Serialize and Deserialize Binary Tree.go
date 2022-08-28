package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:01

 * @Note:	https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
			297. Serialize and Deserialize Binary Tree
			297. 二叉树的序列化与反序列化
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")
	n := len(split)
	if n == 0 {
		return nil
	}
	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if n <= i {
			return nil
		}
		if split[i] == "" {
			i++
			return nil
		}
		root := &TreeNode{}
		root.Val, _ = strconv.Atoi(split[i])
		i++
		root.Left = dfs()
		root.Right = dfs()
		return root
	}
	return dfs()
}
