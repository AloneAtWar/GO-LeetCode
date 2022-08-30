package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/31 1:36

 * @Note:	https://leetcode.cn/problems/validate-stack-sequences/
			946. Validate Stack Sequences
			946. 验证栈序列
 **/

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(popped)
	stack := []int{}
	i := 0
	for j := 0; j < n; j++ {
		for i < n && (len(stack) == 0 || stack[len(stack)-1] != popped[j]) {
			stack = append(stack, pushed[i])
			i++
		}
		if i == n && stack[len(stack)-1] != popped[j] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
