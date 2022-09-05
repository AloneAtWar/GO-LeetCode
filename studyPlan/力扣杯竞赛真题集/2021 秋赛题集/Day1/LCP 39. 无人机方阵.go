package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/4 4:42

 * @Note:	https://leetcode.cn/problems/0jQkd0/?plan=lccup&plan_progress=zn2smkm
			LCP 39. 无人机方阵
 **/

// 错误示范
//func minimumSwitchingTimes(source [][]int, target [][]int) (result int) {
//	table1 := make(map[int]int)
//	table2 := make(map[int]int)
//	for i, ints := range source {
//		for j, v := range ints {
//			table1[v]++
//			table2[target[i][j]]++
//		}
//	}
//	for i, v := range table2 {
//		result+=abs(v-table1[i])
//	}
//	return
//}

func minimumSwitchingTimes(source [][]int, target [][]int) (result int) {
	table := make(map[int]int)
	for i, ints := range source {
		for j, v := range ints {
			table[v]++
			table[target[i][j]]--
		}
	}
	for _, v := range table {
		result += abs(v)
	}
	return result / 2
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
