package main

import "strconv"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 20:22

 * @Note:	https://leetcode.cn/problems/day-of-the-year/
			1154. Day of the Year
			1154. 一年中的第几天
 **/

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		days[1]++
	}

	ans := day
	for _, d := range days[:month-1] {
		ans += d
	}
	return ans
}
