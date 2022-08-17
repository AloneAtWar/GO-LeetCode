package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 14:57

 * @Note:	https://leetcode.cn/problems/remove-interval/
			1272. Remove Interval
			1272. 删除区间
 **/

func removeInterval(intervals [][]int, toBeRemoved []int) (result [][]int) {
	for _, interval := range intervals {
		//    前                                    后
		// 当前终点 早于 要删除区间的起点  或者 当前起点 晚于 删除区间的重点
		if interval[1] <= toBeRemoved[0] || interval[0] >= toBeRemoved[1] {
			result = append(result, interval)
		} else {
			// 前面多一节 或 后面多一节 或者 整个区间都要被删除
			if interval[0] < toBeRemoved[0] {
				result = append(result, []int{interval[0], toBeRemoved[0]})
			}
			if interval[1] > toBeRemoved[1] {
				result = append(result, []int{toBeRemoved[1], interval[1]})
			}
		}
	}
	return
}
