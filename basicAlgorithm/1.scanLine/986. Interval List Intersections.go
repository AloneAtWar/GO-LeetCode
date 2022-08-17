package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 15:51

 * @Note:	https://leetcode.cn/problems/interval-list-intersections/
			986. Interval List Intersections
			986. 区间列表的交集
 **/

func intervalIntersection(firstList [][]int, secondList [][]int) (result [][]int) {
	n, m := len(firstList), len(secondList)
	i, j := 0, 0
	for i != n && j != m {
		left := max(firstList[i][0], secondList[j][0])
		right := min(firstList[i][1], secondList[j][1])
		if left <= right {
			intersection := []int{left, right}
			result = append(result, intersection)
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}

	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
