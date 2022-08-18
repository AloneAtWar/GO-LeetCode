package main

import (
	"fmt"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 4:16

 * @Note:
			1696. Jump Game VI
			1696. 跳跃游戏 VI
 **/

// 第一想法 肯定是用 dp 但这里讲的是 单调队列
func maxResult(nums []int, k int) int {
	n := len(nums)
	var list []int
	// 到达索引点的最大得分
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = nums[i]
		// 一步跳跃不过去
		for len(list) != 0 && i-list[0] > k {
			list = list[1:]
		}
		if len(list) > 0 {
			sum[i] += sum[list[0]]
		}
		// 递减
		for len(list) != 0 && sum[list[len(list)-1]] < sum[i] {
			list = list[:len(list)-1]
		}
		list = append(list, i)
	}
	return sum[n-1]
}

func main() {
	fmt.Println(maxResult([]int{10, -5, -2, 4, 0, 3}, 3))
}
