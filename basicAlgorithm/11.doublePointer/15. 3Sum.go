package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 23:23

 * @Note:	https://leetcode.cn/problems/3sum/
			15. 3Sum
			15. 三数之和
 **/

func threeSum(nums []int) (result [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		first := nums[i]
		if first > 0 {
			break
		}
		left, right := i+1, n-1
		for left < right {
			if left != i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right != n-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			sum := nums[left] + nums[right] + first
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{first, nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return
}
