package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/1 0:29

 * @Note:	https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
			1475. Final Prices With a Special Discount in a Shop
			1475. 商品折扣后的最终价格
 **/

func finalPrices(prices []int) []int {
	n := len(prices)
	result := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = prices[i]
		} else {
			result[i] = prices[i] - stack[len(stack)-1]
		}
		stack = append(stack, prices[i])
	}
	return result
}
