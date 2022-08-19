package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 20:02

 * @Note:	https://leetcode.cn/problems/maximum-gap/
			164. Maximum Gap
			164. 最大间距
 **/

// 桶排序
//虽说很麻烦，但还是再写写看吧
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	m := 1
	max := findMax(nums...)
	size := 10
	temp := make([]int, len(nums))
	for max/m > 0 {
		barrel := make([]int, size)
		for _, num := range nums {
			barrel[(num/m)%size]++
		}
		for i := 1; i < size; i++ {
			barrel[i] += barrel[i-1]
		}
		for i := len(nums) - 1; i >= 0; i-- {
			barrel[(nums[i]/m)%size]--
			index := barrel[(nums[i]/m)%size]
			temp[index] = nums[i]
		}
		temp, nums = nums, temp
		m *= 10
	}
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > result {
			result = nums[i] - nums[i-1]
		}
	}
	return result
}

func findMax(ints ...int) int {
	result := ints[0]
	for _, v := range ints[1:] {
		if v > result {
			result = v
		}
	}
	return result
}
