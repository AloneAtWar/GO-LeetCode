package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 19:53

 * @Note:	https://leetcode.cn/problems/trapping-rain-water/
			42. Trapping Rain Water
			42. 接雨水
 **/

// 这题老经典了  先来手简单的 dp  竖着来算雨的面积
//func trap(height []int) (result int) {
//	n:=len(height)
//	dpLeft:=make([]int,n)
//	dpRight:=make([]int,n)
//	dpLeft[0]=height[0]
//	for i := 1; i <n ; i++ {
//		dpLeft[i]=max(dpLeft[i-1],height[i])
//	}
//	dpRight[n-1]=height[n-1]
//	for j := n-2; j >=0 ; j-- {
//		dpRight[j]=max(dpRight[j+1],height[j])
//	}
//	for i := 1; i < n-1; i++ {
//		result+=min(dpRight[i],dpLeft[i])-height[i]
//	}
//	return
//}
//
//func max(i,j int)int{
//	if i>j{
//		return i
//	}else{
//		return j
//	}
//}
//
//func min(a,b int)int{
//	if a>b{
//		return b
//	}
//	return a
//}

// 正经的单调栈的算法
// 利用的特性是说 若要积水 则前面至少有两个高度 呈递减状
// 横着算
func trap(height []int) (result int) {
	var stack []int
	for i, h := range height {
		for len(stack) != 0 && height[stack[len(stack)-1]] < h {
			if len(stack) >= 2 {
				result += (i - stack[len(stack)-2] - 1) * (min(h, height[stack[len(stack)-2]]) - height[stack[len(stack)-1]])
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
