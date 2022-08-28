package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:11

 * @Note:	https://leetcode.cn/problems/serialize-and-deserialize-n-ary-tree/
			428. Serialize and Deserialize N-ary Tree
			428. 序列化和反序列化 N 叉树
 **/
type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	n := len(root.Children)
	builder := strings.Builder{}
	result := fmt.Sprintf("%d,%d", root.Val, n)
	builder.WriteString(result)
	for i := 0; i < n; i++ {
		builder.WriteString("," + this.serialize(root.Children[i]))
	}
	return builder.String()
}

func (this *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	list := strings.Split(data, ",")
	n := len(list)
	i := 0
	var dfs func() *Node
	dfs = func() *Node {
		if i >= n {
			return nil
		}
		val, _ := strconv.Atoi(list[i])
		length, _ := strconv.Atoi(list[i+1])
		i += 2
		result := &Node{Val: val}
		result.Children = make([]*Node, length)
		for i := 0; i < length; i++ {
			result.Children[i] = dfs()
		}
		return result
	}
	return dfs()
}
