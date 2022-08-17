package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date: 	2022/8/17 13:50

 * @Note:	https://leetcode.cn/problems/meeting-rooms/
			252. Meeting Rooms
			252. 会议室
 **/

// 把每场会议按照开始时间进行排序 只需要保证说 当前会议的开始时间不要早于上一场会议的结束时间
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pre := -1
	for _, interval := range intervals {
		if interval[0] < pre {
			return false
		}
		pre = interval[1]
	}
	return true
}
