package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 0:05

 * @Note:

 **/
// 其实本质上 就是 正序遍历 之后将结果倒置了
func postorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	// 指的是 还没有把左右加进栈中
	var stack []*TreeNode = []*TreeNode{root}
	// 把左右加进栈中了
	var stack2 []*TreeNode
	for len(stack) != 0 {
		out := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if out.Left != nil {
			stack = append(stack, out.Left)
		}
		if out.Right != nil {
			stack = append(stack, out.Right)
		}
		stack2 = append(stack2, out)
	}
	for len(stack2) != 0 {
		out := stack2[len(stack2)-1]
		stack = stack2[:len(stack2)-1]
		result = append(result, out.Val)
	}
	return
}
