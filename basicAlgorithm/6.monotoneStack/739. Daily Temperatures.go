package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 19:02

 * @Note:	https://leetcode.cn/problems/daily-temperatures/
			739. Daily Temperatures
			739. 每日温度
 **/

func dailyTemperatures(temperatures []int) (result []int) {
	var stack []int
	result = make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return
}
