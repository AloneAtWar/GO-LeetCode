package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 17:41

 * @Note:	https://leetcode.cn/problems/meeting-scheduler/
			1229. Meeting Scheduler
			1229. 安排会议日程
 **/

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {

	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})
	n, m := len(slots1), len(slots2)
	i, j := 0, 0
	for i != n && j != m {
		left := max(slots1[i][0], slots2[j][0])
		right := min(slots1[i][1], slots2[j][1])
		if left <= right && right-left >= duration {
			//有交集
			return []int{left, left + duration}
		}
		if slots1[i][1] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}
	return []int{}
}
