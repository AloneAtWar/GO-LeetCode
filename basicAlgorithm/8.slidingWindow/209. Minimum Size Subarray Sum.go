package main

import "math"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 15:59

 * @Note:

 **/

func minSubArrayLen(target int, nums []int) (result int) {
	n := len(nums)
	result = math.MaxInt64
	left := 0
	currSum := 0
	for i := 0; i < n; i++ {
		currSum += nums[i]
		for currSum >= target {
			if i-left+1 < result {
				result = i - left + 1
			}
			currSum -= nums[left]
			left++
		}
	}
	if result == math.MaxInt64 {
		return 0
	} else {
		return result
	}
}
