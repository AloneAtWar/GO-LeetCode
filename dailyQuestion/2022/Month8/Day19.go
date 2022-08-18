package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 1:53

 * @Note:	https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time/
			1450. Number of Students Doing Homework at a Given Time
			1450. 在既定时间做作业的学生人数
 **/

// 妥妥的送分题
func busyStudent(startTime []int, endTime []int, queryTime int) (result int) {
	n := len(startTime)
	for i := 0; i < n; i++ {
		if queryTime >= startTime[i] && endTime[i] >= queryTime {
			result++
		}
	}
	return
}
