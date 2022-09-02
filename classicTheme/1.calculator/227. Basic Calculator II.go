package main

import "strings"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 15:09

 * @Note:	https://leetcode.cn/problems/basic-calculator-ii/
			227. Basic Calculator II
			227. 基本计算器 II
 **/

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	var stack []int
	n := len(s)
	operator := '+'
	num := 0
	for i := 0; i < n; i++ {
		if s[i] <= '9' && s[i] >= '0' {
			num = num*10 + int(s[i]-'0')
			if i != n-1 {
				continue
			}
		}
		if operator == '+' {
			stack = append(stack, num)
		} else if operator == '-' {
			stack = append(stack, -num)
		} else if operator == '*' {
			out := stack[len(stack)-1]
			stack[len(stack)-1] = num * out
		} else {
			out := stack[len(stack)-1]
			stack[len(stack)-1] = out / num
		}
		num = 0
		operator = int32(s[i])
	}
	for _, v := range stack {
		num += v
	}
	return num
}
