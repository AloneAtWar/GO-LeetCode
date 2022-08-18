package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 16:45

 * @Note:	https://leetcode.cn/problems/next-greater-element-i/
			496. Next Greater Element I
			496. 下一个更大元素 I
 **/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	table := map[int]int{}
	for i, num := range nums2 {
		table[num] = i
	}
	temp := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			temp[i] = -1
		} else {
			temp[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	result := make([]int, 0, len(nums1))
	for _, i := range nums1 {
		result = append(result, temp[table[i]])
	}
	return result
}
