package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 15:33

 * @Note:	https://leetcode.cn/problems/remove-covered-intervals/
			1288. Remove Covered Intervals
			1288. 删除被覆盖区间
 **/

func removeCoveredIntervals(intervals [][]int) (result int) {
	sort.Slice(intervals, func(i, j int) bool {
		// 能覆盖后者的放在前面
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	result++
	pre := intervals[0]
	for _, interval := range intervals {
		if interval[0] >= pre[0] && pre[1] >= interval[1] {

		} else {
			result++
			pre = interval
		}
	}
	return
}
