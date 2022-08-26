package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/27 2:37

 * @Note:	https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/
			1464. Maximum Product of Two Elements in an Array
			1464. 数组中两元素的最大乘积
			妥妥的送分题啊！！！！！！！！！！
 **/

func maxProduct(nums []int) int {
	One, two := nums[0], nums[1]
	if One < two {
		One, two = two, One
	}
	for _, v := range nums[2:] {
		if v > One {
			One, two = v, One
		} else if v > two {
			two = v
		}
	}
	return (One - 1) * (two - 1)
}
