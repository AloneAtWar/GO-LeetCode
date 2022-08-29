package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 23:13

 * @Note:	https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
			1123. Lowest Common Ancestor of Deepest Leaves
			1123. 最深叶节点的最近公共祖先
 **/

// 用来测试1676
//func lcaDeepestLeaves(root *TreeNode) *TreeNode {
//	if root==nil{
//		return root
//	}
//	list:=[]*TreeNode{root}
//	next:=[]*TreeNode{}
//	for len(list)!=0{
//		for _, node := range list {
//			if node.Left!=nil{
//				next=append(next,node.Left)
//			}
//			if node.Right!=nil{
//				next=append(next,node.Right)
//			}
//		}
//		if len(next)==0{
//			return lowestCommonAncestor1676(root,list)
//		}
//		list,next=next,[]*TreeNode{}
//	}
//	return nil
//}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var getDeep func(root *TreeNode, deep int) (*TreeNode, int)
	getDeep = func(root *TreeNode, deep int) (*TreeNode, int) {
		if root == nil {
			return nil, deep
		}
		node1, deep1 := getDeep(root.Left, deep+1)
		node2, deep2 := getDeep(root.Right, deep+1)
		if deep1 == deep2 {
			return root, deep1
		} else if deep1 > deep2 {
			return node1, deep1
		} else {
			return node2, deep2
		}
	}
	result, _ := getDeep(root, 0)
	return result
}
