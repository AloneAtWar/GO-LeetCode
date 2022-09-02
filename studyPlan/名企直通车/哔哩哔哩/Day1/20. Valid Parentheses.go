package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 20:03

 * @Note:	https://leetcode.cn/problems/valid-parentheses/?plan=zhitongche&plan_progress=zmq4dzv
			20. Valid Parentheses
			20. 有效的括号
 **/

func isValid(s string) bool {
	var stack []int32
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if v == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if v == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
