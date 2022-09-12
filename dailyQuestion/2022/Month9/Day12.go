package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/12 9:02

 * @Note:	https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/
			1608. Special Array With X Elements Greater Than or Equal X
			1608. 特殊数组的特征值
 **/

//func specialArray(nums []int) int {
//	n:=len(nums)
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i]<nums[j]
//	})
//	left,right:=0,nums[n-1]
//	check:=func(checkNum int)int{
//		l,r:=0,n-1
//		for l<r{
//			m:=(l+r+1)>>1
//			if nums[m]>=checkNum{
//				r=m-1
//			}else{
//				l=m
//			}
//		}
//		if nums[r]<checkNum{
//			r++
//		}
//		return n-r
//	}
//	for left<=right{
//		mid:=(left+right)>>1
//		this:=check(mid)
//		if this==mid{
//			return mid
//		}else if this>mid{
//			left=mid+1
//		}else{
//			right=mid-1
//		}
//	}
//	return -1
//}
func specialArray(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

func main() {
	specialArray([]int{0, 4, 3, 0, 4})
}
