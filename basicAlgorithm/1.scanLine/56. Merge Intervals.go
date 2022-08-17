package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date: 	2022/8/17 14:21

 * @Note:	https://leetcode.cn/problems/merge-intervals/
			56. Merge Intervals
			56. 合并取键
 **/

//  把区间按照  开始时间进行排序 若开始时间相同 则将范围更大的区间放在更前头
//  每次取出来一个区间 判断这个取键会不会跟前一个区间有冲突 如果没有 那么前一个区间就可以放在result里
//  因为该区间不会对后面的区间有其他的影响 如果有冲突 那么当前区间就要与上一个区间做合并
//  合并的话 因为上一个区间的起始时间一定小于等于当前区间 所以就只合并 结束时间
//  合并之后依然不能放入 result 因为 他也有可能跟下下个区间有冲突

func merge(intervals [][]int) (result [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	pre := intervals[0]
	for _, ints := range intervals[1:] {
		if ints[0] > pre[1] {
			result = append(result, pre)
			pre = ints
		} else {
			pre[1] = max(pre[1], ints[1])
		}
	}
	result = append(result, pre)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
