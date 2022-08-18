package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 16:53

 * @Note:	https://leetcode.cn/problems/next-greater-element-ii/
			503. Next Greater Element II
			503. 下一个更大元素 II
 **/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	temp := make([]int, 2*len(nums))
	var stack []int
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			temp[i] = -1
		} else {
			temp[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}
	return temp[:n]
}
