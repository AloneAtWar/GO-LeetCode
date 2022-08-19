package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 16:22

 * @Note:	https://leetcode.cn/problems/subarrays-with-k-different-integers/
			992. Subarrays with K Different Integers
			992. K 个不同整数的子数组
 **/

func subarraysWithKDistinct(nums []int, k int) (result int) {
	most := func(k int) int {
		table := make(map[int]int)
		result := 0
		left := 0
		for i, num := range nums {
			table[num]++
			for len(table) > k {
				table[nums[left]]--
				if table[nums[left]] == 0 {
					delete(table, nums[left])
				}
				left++
			}
			result += i - left + 1
		}
		return result
	}
	return most(k) - most(k-1)
}
