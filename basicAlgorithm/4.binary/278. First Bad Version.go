package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 14:08

 * @Note:	https://leetcode.cn/problems/first-bad-version/
			278. First Bad Version
			278. 第一个错误的版本
 **/
//编译
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) >> 1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
