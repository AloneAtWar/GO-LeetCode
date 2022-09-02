package main

import "strings"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 15:33

 * @Note:	https://leetcode.cn/problems/basic-calculator-iii/
			772. Basic Calculator III
			772. 基本计算器 III
 **/

func calculate(s string) int {
	var i = 0
	s = strings.Replace(s, " ", "", -1)
	n := len(s)
	var f func() int
	f = func() int {
		var stack []int
		operator := '+'
		num := 0
		for ; i < n; i++ {
			if s[i] <= '9' && s[i] >= '0' {
				num = num*10 + int(s[i]-'0')
				if i != n-1 {
					continue
				}
			}
			if s[i] == '(' {
				i++
				num = f()
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
			if s[i] == ')' {

				break
			}
		}
		for _, v := range stack {
			num += v
		}
		return num
	}
	return f()
}
