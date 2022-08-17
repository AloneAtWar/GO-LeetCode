package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date: 	2022/8/17 13:47

 * @Note:	https://www.lintcode.com/problem/391/
			391.Number of Airplanes in the Sky（LintCode)
			391.数飞机
 **/

type Interval struct {
	Start, End int
}

// 方法一： 只关注飞机的起飞时间和降落时间
// 起飞时间影响当前天上的飞机个数+1
// 降落时间影响当前天上的飞机个数-1
// 算出飞机时间最大以及最小的时候

func CountOfAirplanes(airplanes []*Interval) int {
	type moments struct {
		t int
		s int // 0 代表起飞 1 代表降落
	}
	lists := make([]*moments, 0, 2*len(airplanes))
	for _, airplane := range airplanes {
		s := &moments{t: airplane.Start, s: 0}
		e := &moments{t: airplane.End, s: 1}
		lists = append(lists, s, e)
	}
	sort.Slice(lists, func(i, j int) bool {
		if lists[i].t != lists[j].t {
			return lists[i].t < lists[j].t
		}
		//应该先降落再起飞
		return lists[i].s == 1
	})
	curr := 0
	result := 0
	for _, list := range lists {
		if list.s == 0 {
			curr++
		} else {
			curr--
		}
		result = max(result, curr)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
