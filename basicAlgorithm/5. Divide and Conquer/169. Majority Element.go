package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 15:35

 * @Note:	https://leetcode.cn/problems/majority-element/
			169. Majority Element
			169. 多数元素
 **/

// 投票法
// 这题用分治法太傻了
func majorityElement(nums []int) int {
	currNum := nums[0]
	size := 1
	for _, num := range nums[1:] {
		if num == currNum {
			size++
		} else {
			size--
			if size == 0 {
				size++
				currNum = num
			}
		}
	}
	return currNum
}
