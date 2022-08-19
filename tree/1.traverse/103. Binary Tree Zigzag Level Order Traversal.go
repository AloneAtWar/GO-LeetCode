package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 1:31

 * @Note:

 **/

//func zigzagLevelOrder(root *TreeNode) (result [][]int) {
//	if root==nil{
//		return nil
//	}
//	pre:=[]*TreeNode{root}
//	var next []*TreeNode
//	k:=0
//	for len(pre)!=0{
//		var this []int
//		if k%2==0{
//			for _, node := range pre {
//				this = append(this, node.Val)
//				if node.Left != nil {
//					next = append(next, node.Left)
//				}
//				if node.Right != nil {
//					next = append(next, node.Right)
//				}
//			}
//		}else{
//			for i := len(pre)-1; i >=0; i-- {
//				node:=pre[i]
//				if node.Left != nil {
//					next = append(next, node.Left)
//				}
//				if node.Right != nil {
//					next = append(next, node.Right)
//				}
//			}
//		}
//		k++
//		result=append(result,this)
//		pre=next
//		next=[]*TreeNode{}
//	}
//	return
//}

// 再用到 bfs 的时候 容易出错的点是说 想着 结果倒叙 然后也把增加队列的顺序也更换了 所以导致结果 错误
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return
}
