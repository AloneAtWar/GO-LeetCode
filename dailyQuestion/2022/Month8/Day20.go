package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/20 0:32

 * @Note:

 **/

// 想到了说 需要知道每个节点 某一个范围的最大值
// 但是想的dp 跟直接递归没有什么两样
//func constructMaximumBinaryTree(nums []int) *TreeNode {
//	n:=len(nums)
//	dpMax:=make([][]int,n)
//	for i := 0; i < n; i++ {
//		dpMax[i]=make([]int,n)
//		dpMax[i][i]=i
//	}
//	for i := 0; i < n; i++ {
//		for j := i+1; j < n; j++ {
//			if nums[dpMax[i][j-1]]<nums[j]{
//				dpMax[i][j]=nums[j]
//			}else{
//				dpMax[i][j]=dpMax[i][j-1]
//			}
//		}
//	}
//	return build(nums,dpMax,0,n-1)
//}
//
//func build(nums []int, dpMax [][]int,left int, right int) *TreeNode {
//	if left>right{
//		return nil
//	}else if left==right{
//		return &TreeNode{Val: nums[left]}
//	}
//	return &TreeNode{Val: nums[dpMax[left][right]],Left: build(nums,dpMax,left,dpMax[left][right]-1),Right: build(nums,dpMax,dpMax[left][right]+1,right)}
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 借由提示 单调栈 自己实现一下
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)

	var stack1 []int
	leftMax := make([]int, n)

	var stack2 []int
	rightMax := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack1) != 0 && nums[stack1[len(stack1)-1]] < nums[i] {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack1) == 0 {
			leftMax[i] = -1
		} else {
			leftMax[i] = stack1[len(stack1)-1]
		}
		stack1 = append(stack1, i)
		for len(stack2) != 0 && nums[stack2[len(stack2)-1]] < nums[n-1-i] {
			stack2 = stack2[:len(stack2)-1]
		}
		if len(stack2) == 0 {
			rightMax[n-i-1] = -1
		} else {
			rightMax[n-i-1] = stack2[len(stack2)-1]
		}
		stack2 = append(stack2, n-i-1)
	}
	return build(nums, leftMax, rightMax, leftMax, 0, 0, len(nums)-1)
}

func build(nums []int, stack1 []int, stack2 []int, use []int, begin int, left int, right int) *TreeNode {
	if left > right {
		return nil
	} else if left == right {
		return &TreeNode{Val: nums[left]}
	}
	i := begin
	// 若 i==-1 或者 i 不在范围之内 都说明是 它为区域中最大的
	for left <= use[i] && use[i] >= right {
		i = use[i]
	}
	return &TreeNode{Val: nums[i], Left: build(nums, stack1, stack2, stack1, i-1, left, i-1), Right: build(nums, stack1, stack2, stack2, i+1, i+1, right)}
}

//func constructMaximumBinaryTree(nums []int) *TreeNode {
//    tree := make([]*TreeNode, len(nums))
//    stk := []int{}
//    for i, num := range nums {
//        tree[i] = &TreeNode{Val: num}
//        for len(stk) > 0 && num > nums[stk[len(stk)-1]] {
//            tree[i].Left = tree[stk[len(stk)-1]]
//            stk = stk[:len(stk)-1]
//        }
//        if len(stk) > 0 {
//            tree[stk[len(stk)-1]].Right = tree[i]
//        }
//        stk = append(stk, i)
//    }
//    return tree[stk[0]]
//}
//后面这个就太绝了把
//在最终构造出的树上，以 \textit{nums}[i]nums[i] 为根节点的子树，在原数组中对应的区间，左边界为 \textit{nums}[i]nums[i] 左侧第一个比它大的元素所在的位置，右边界为 \textit{nums}[i]nums[i] 右侧第一个比它大的元素所在的位置。左右边界均为开边界。
//
//如果某一侧边界不存在，则那一侧边界为数组的边界。如果两侧边界均不存在，说明其为最大值，即根节点。
//
// 只能说牛逼！
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/maximum-binary-tree/solution/zui-da-er-cha-shu-by-leetcode-solution-lbeo/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
