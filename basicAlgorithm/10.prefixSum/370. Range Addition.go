package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 22:50

 * @Note:	https://leetcode.cn/problems/range-addition/
			370. Range Addition
			370. 区间加法
 **/

// 延迟增加  前缀和
func getModifiedArray(length int, updates [][]int) []int {
	result := make([]int, length)
	for _, update := range updates {
		start, end, inc := update[0], update[1], update[2]
		result[start] += inc
		if end != length-1 {
			result[end+1] -= inc
		}
	}
	for i := 1; i < length; i++ {
		result[i] += result[i-1]
	}
	return result
}
