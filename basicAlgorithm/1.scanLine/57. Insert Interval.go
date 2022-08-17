package main

/**

 * @Author: AloneAtWar

 * @Date: 	2022/8/17 14:39

 * @Note:	https://leetcode.cn/problems/insert-interval/
			57. Insert Interval
			57. 插入区间
 **/

// 	一个简单的方法室 直接 插入 intervals 之后再 合并区间
// 	这里我就不过多地再进行演示
//  挨个判断 新区间是否会跟已有区间冲突
func insert(intervals [][]int, newInterval []int) (result [][]int) {
	for _, interval := range intervals {
		// 如果不增加新区间 或者 增加的新区间不冲突
		if newInterval == nil || newInterval[0] > interval[1] {
			result = append(result, interval)
		} else if newInterval[1] < interval[0] {
			result = append(result, newInterval, interval)
			newInterval = nil
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	if newInterval != nil {
		result = append(result, newInterval)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
