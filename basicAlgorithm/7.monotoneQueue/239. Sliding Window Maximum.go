package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 2:01

 * @Note:	https://leetcode.cn/problems/sliding-window-maximum/
			239. Sliding Window Maximum
			239. 滑动窗口最大值
 **/

func maxSlidingWindow(nums []int, k int) (result []int) {
	var list []int
	for i, num := range nums {
		for len(list) != 0 && i-list[0] >= k {
			list = list[1:]
		}
		for len(list) != 0 && nums[list[len(list)-1]] <= num {
			list = list[:len(list)-1]
		}
		list = append(list, i)
		if i >= k-1 {
			result = append(result, nums[list[0]])
		}
	}
	return result
}
