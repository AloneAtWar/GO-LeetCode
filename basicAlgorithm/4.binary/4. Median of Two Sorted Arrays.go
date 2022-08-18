package main

import (
	"fmt"
	"math"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 14:10

 * @Note:	https://leetcode.cn/problems/median-of-two-sorted-arrays/
			4. Median of Two Sorted Arrays
			4. 寻找两个正序数组的中位数
 **/

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{1}))
}

// num1[]
// num2[]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	if m > n {
		// 让 n > m
		nums1, nums2, n, m = nums2, nums1, m, n
	}
	left, right := 0, m
	for left <= right {
		// 下边界为 mid
		lowerBoundary := (left + right) >> 1
		// 上边界
		upperBoundary := (m+n)>>1 - lowerBoundary
		upperBoundaryL := find(nums1, upperBoundary, 0)
		upperBoundaryR := find(nums1, upperBoundary, 1)
		lowerBoundaryL := find(nums2, lowerBoundary, 0)
		lowerBoundaryR := find(nums2, lowerBoundary, 1)
		if upperBoundaryR >= lowerBoundaryL && lowerBoundaryR >= upperBoundaryL {
			if (m+n)%2 == 0 {
				return float64(min(upperBoundaryR, lowerBoundaryR)+max(upperBoundaryL, lowerBoundaryL)) / 2
			} else {
				// 说明 中位数 就是一个数 切 左边
				// 左半边的个数一定会小于右半边的个数 所有  若可能右侧最小的那个一个就为中位数
				return float64(min(upperBoundaryR, lowerBoundaryR))
			}
		} else if upperBoundaryR < lowerBoundaryL {
			right = lowerBoundary - 1
		} else {
			left = lowerBoundary + 1
		}
		// 说明 中位数是 两个数之和 /2

	}
	return 0
}
func find(nums []int, index, t int) int {
	// 左侧
	if t == 0 {
		if index == 0 {
			return math.MinInt64
		} else {
			return nums[index-1]
		}
		//右侧
	} else {
		if index == len(nums) {
			return math.MaxInt64
		} else {
			return nums[index]
		}
	}
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
