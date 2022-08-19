package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 15:13

 * @Note:	https://leetcode.cn/problems/kth-largest-element-in-an-array/
			215. Kth Largest Element in an Array
			215. 数组中的第K个最大元素
 **/

// 快速选择算法
func findKthLargest(nums []int, k int) int {
	return choose(nums, 0, len(nums)-1, len(nums)-k)
}

// 这里的对应关系写的可能有点不太好
func choose(nums []int, left int, right int, k int) int {
	index := left
	indexNum := nums[index]
	small := left + 1
	for i := left + 1; i <= right; i++ {
		if nums[i] < indexNum {
			nums[small], nums[i] = nums[i], nums[small]
			small++
		}
	}
	nums[index], nums[small-1] = nums[small-1], nums[index]
	if small-1 == k {
		return indexNum
	} else if small-1 > k {
		return choose(nums, left, small-2, k)
	} else {
		return choose(nums, small, right, k)
	}
}
