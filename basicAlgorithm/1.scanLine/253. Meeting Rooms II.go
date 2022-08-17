package main

import (
	"container/heap"
	"sort"
)

/**

 * @Author: AloneAtWar

 * @Date: 	2022/8/17 13:55

 * @Note:	https://leetcode.cn/problems/meeting-rooms-ii/
			253. Meeting Rooms II.go
			253. 会议室 II
 **/

// 先将会议按照 会议的开始时间进行排序
// 利用堆的性质 最小堆
// 每次有一个新的会议时 检查一下能否用已有的会议室进行安排 若不能则需要开新的会议室
// 记录所开会议室最多的数量

// 实现 Heap 接口
type meetingHeap [][]int

func (m *meetingHeap) Len() int {
	return len(*m)
}

func (m *meetingHeap) Less(i, j int) bool {
	return (*m)[i][1] < (*m)[j][1]
}

func (m *meetingHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *meetingHeap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *meetingHeap) Pop() interface{} {
	out := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return out
}

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &meetingHeap{intervals[0]}
	result := 1
	for _, ints := range intervals[1:] {
		// 如果当前会议 开始时间 早于 所有 会议室的结束时间 那么就需要重新开一个新会议室
		if ints[0] < (*h)[0][1] {
			heap.Push(h, ints)
			result = max(result, h.Len())
		} else {
			out := heap.Pop(h).([]int)
			out[1] = max(out[1], ints[1])
			heap.Push(h, ints)
		}
	}
	return result
}

func main() {
	i := [][]int{{1, 8}, {6, 20}, {9, 16}, {13, 17}}
	minMeetingRooms(i)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
