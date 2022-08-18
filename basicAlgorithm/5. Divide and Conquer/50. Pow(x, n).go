package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 15:47

 * @Note:	https://leetcode.cn/problems/powx-n/
			50. Pow(x, n)
			50. Pow(x, n)
 **/

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}
	curr := 1.0
	this := x
	for n != 0 {
		if n&1 != 0 {
			curr *= this
		}

		this *= this

		n = n >> 1
	}
	return curr
}
