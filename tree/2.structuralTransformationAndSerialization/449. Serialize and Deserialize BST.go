package main

import (
	"math"
	"strconv"
	"strings"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 0:21

 * @Note:

 **/

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
	builder := strings.Builder{}
	builder.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil {
		builder.WriteString(",")
		builder.WriteString(this.serialize(root.Left))
	}
	if root.Right != nil {
		builder.WriteString(",")
		builder.WriteString(this.serialize(root.Right))
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	datas := strings.Split(data, ",")
	n := len(datas)
	i := 0
	var dfs func(max, min int) *TreeNode
	dfs = func(max, min int) *TreeNode {
		if i >= n {
			return nil
		}
		v, _ := strconv.Atoi(datas[i])
		if v > min && v < max {
			root := &TreeNode{}
			root.Val = v
			i++
			root.Left = dfs(v, min)
			root.Right = dfs(max, v)
			return root
		}
		return nil
	}
	return dfs(math.MaxInt64, math.MinInt)
}
