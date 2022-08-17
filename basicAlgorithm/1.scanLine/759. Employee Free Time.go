package main

import (
	"math"
	"sort"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 16:26

 * @Note:	https://leetcode.cn/problems/employee-free-time/
			759. Employee Free Time
			759. 员工空闲时间
 **/

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) (result []*Interval) {
	intervals := []*Interval{}
	for _, i := range schedule {
		for _, interval := range i {
			intervals = append(intervals, interval)
		}
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End > intervals[i].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	left := intervals[0].End
	for _, interval := range intervals[1:] {
		if interval.Start > left {
			result = append(result, &Interval{Start: left, End: interval.Start})
		}
		left = max(left, interval.End)
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
