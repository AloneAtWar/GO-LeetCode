package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 19:44

 * @Note:	https://leetcode.cn/problems/remove-k-digits/
			402. Remove K Digits.go
			402. 移掉 K 位数字
 **/

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack []byte
	for _, v := range num {
		for k > 0 && len(stack) != 0 && byte(v) < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, byte(v))
	}
	stack = stack[:len(stack)-k]
	left := 0
	for left < len(stack) {
		if stack[left] == '0' {
			left++
		} else {
			break
		}
	}
	stack = stack[left:]
	if len(stack) == 0 {
		return "0"
	} else {
		return string(stack)
	}
}
