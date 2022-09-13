package main

import "strconv"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/13 9:32

 * @Note:	https://leetcode.cn/problems/maximum-swap/
			670. Maximum Swap
			670. 最大交换
 **/

func maximumSwap(num int) int {
	numStr := []byte(strconv.Itoa(num))
	n := len(numStr)
	index := -1
	for i := n - 1; i > 0; i-- {
		if numStr[i] > numStr[i-1] && (index == -1 || numStr[index] < numStr[i]) {
			index = i
		}
	}
	if index == -1 {
		return num
	}
	for i := n - 1; i > index; i-- {
		if numStr[i] == numStr[index] {
			index = i
			break
		}
	}
	for i := 0; i < n; i++ {
		if numStr[i] < numStr[index] {
			numStr[index], numStr[i] = numStr[i], numStr[index]
			break
		}
	}
	result, _ := strconv.Atoi(string(numStr))
	return result
}
