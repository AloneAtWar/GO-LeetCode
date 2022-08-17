package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 15:14

 * @Note:	https://leetcode.cn/problems/non-overlapping-intervals/
			435. Non-overlapping Intervals

 **/

// 难点在于 考虑删除那个区间 若按照开始时间进行排序 则应该删除结束时间更长的那个
// 同理 若按照结束时间进行排序 则应该删除开时间更早的那个
func eraseOverlapIntervals(intervals [][]int) (result int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pre := intervals[0]
	for _, ints := range intervals[1:] {
		if ints[0] < pre[1] {
			result++
			if pre[1] > ints[1] {
				pre = ints
			}
		} else {
			pre = ints
		}
	}
	return
}
