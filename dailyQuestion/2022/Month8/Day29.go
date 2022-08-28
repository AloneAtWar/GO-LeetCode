package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/29 1:36

 * @Note:	https://leetcode.cn/problems/shuffle-the-array/submissions/
			1470. 重新排列数组
			1470. Shuffle the Array
 **/

func shuffle(nums []int, n int) []int {
	ans := make([]int, n*2)
	for i, num := range nums[:n] {
		ans[2*i] = num
		ans[2*i+1] = nums[n+i]
	}
	return ans
}
