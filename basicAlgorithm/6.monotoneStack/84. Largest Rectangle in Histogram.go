package main

import (
	"fmt"
	"math"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 20:42

 * @Note:	https://leetcode.cn/problems/largest-rectangle-in-histogram/
			84. Largest Rectangle in Histogram
			84. 柱状图中最大的矩形
 **/

// 这个思路有点牛逼 分治的思路
//func largestRectangleArea(heights []int) (result int) {
//	left,right:=0,len(heights)-1
//	return find(heights,left,right)
//}
//
//func find(heights []int, left int, right int) int {
//	if left>right{
//		return 0
//	}
//	if left==right{
//		return heights[left]
//	}
//	minIndex:=left
//	low:=heights[left]
//	for i := left+1; i <= right; i++ {
//		if low>heights[i]{
//			low=heights[i]
//			minIndex=i
//		}
//	}
//	return max(low*(right-left+1),find(heights,left,minIndex-1),find(heights,minIndex+1,right))
//}
//
//func max(nums ...int)int{
//	m:=nums[0]
//	for _, num := range nums {
//		if num>m{
//			m=num
//		}
//	}
//	return m
//}

// 好难理解啊！！！！！！！！！！！
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

//每个节点 只需要知道 他的左边 第一个比自己低 他索引是多少
// 所以联想到到使用递增站
// 递增站 在遇到新元素 比栈顶元素小的时候
// 弹出栈顶元素
// 从遍历到的节点索引 i 到 比弹出元素第一个小的元素 都可以作为 高度为 heights的width
func largestRectangleArea(heights []int) int {
	result := math.MinInt64
	var stack []int
	n := len(heights)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > heights[i] {
			outIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			result = max(result, width*heights[outIndex])
		}
		stack = append(stack, i)
	}
	//当便利完之后 发现还有元素
	// 潜台词是说该节点之后遇不到比该节点小的元素
	// 说明该节点的右边都可以作为 该节点 hegin 的 width
	for len(stack) != 0 {
		outIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var width int
		if len(stack) == 0 {
			width = len(heights)
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		result = max(result, width*heights[outIndex])
	}

	return result
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
