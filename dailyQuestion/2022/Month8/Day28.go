package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/28 19:06

 * @Note:	https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/
			793. Preimage Size of Factorial Zeroes Function
			793. 阶乘函数后 K 个零
 **/

func preimageSizeFZF(k int) int {
	num := 0
	//先确定右界
	left1, right1 := 0, 5*(k+1)-1
	for left1 < right1 {
		mid := (left1 + right1 + 1) >> 1
		num = trailingZeroes(mid)
		if num > k {
			right1 = mid - 1
		} else {
			left1 = mid
		}
	}
	if trailingZeroes(left1) != k {
		return 0
	}
	//再确定左
	left2, right2 := 0, right1
	for left2 < right2 {
		mid := (left2 + right2) >> 1
		num = trailingZeroes(mid)
		if num >= k {
			right2 = mid
		} else {
			left2 = mid + 1
		}
	}
	return left1 - left2 + 1
}

// 先导题
//	172. Factorial Trailing Zeroes
//	172. 阶乘后的零

func trailingZeroes(n int) (result int) {
	for n > 0 {
		n /= 5
		result += n
	}
	return
}
