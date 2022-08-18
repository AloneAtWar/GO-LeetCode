package main

import "fmt"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 2:59

 * @Note:	https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
			1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
			1438. 绝对差不超过限制的最长连续子数组
 **/

// 暴力解法
//func longestSubarray(nums []int, limit int) (result int) {
//	n:=len(nums)
//	for i := 0; i < n; i++ {
//		left,right:=i,i
//		currNum:=nums[i]
//		for left>0 && nums[left-1]>=currNum &&nums[left-1]-currNum<=limit{
//			left--
//		}
//		for right<n-1 && nums[right+1]<=currNum && nums[right+1]-currNum<=limit{
//			right++
//		}
//		currSize:=right-left+1
//		if currSize>result{
//			result=currSize
//		}
//	}
//	return
//}

func longestSubarray(nums []int, limit int) (result int) {
	var increment []int
	var decreas []int
	left := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for len(increment) != 0 && nums[increment[len(increment)-1]] >= nums[i] {
			increment = increment[:len(increment)-1]
		}
		for len(decreas) != 0 && nums[decreas[len(decreas)-1]] <= nums[i] {
			decreas = decreas[:len(decreas)-1]
		}
		increment = append(increment, i)
		decreas = append(decreas, i)
		for nums[decreas[0]]-nums[increment[0]] > limit {
			if decreas[0] == left {
				decreas = decreas[1:]
			}
			if increment[0] == left {
				increment = increment[1:]
			}
			left++
		}
		if i-left+1 > result {
			result = i - left + 1
		}
	}
	return
}

func main() {
	fmt.Println(longestSubarray([]int{2, 4, 1, 3}, 2))
}
