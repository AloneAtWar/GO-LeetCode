package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 14:38

 * @Note:	https://leetcode.cn/problems/split-array-largest-sum/
			410. Split Array Largest Sum
			410. 分割数组的最大值
 **/

func splitArray(nums []int, m int) int {
	left, right := max(nums...), sum(nums...)
	for left < right {
		mid := (left + right) >> 1
		if check(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(nums []int, checkNum int, m int) bool {
	i := 1
	curr := 0
	for _, num := range nums {
		curr += num
		if curr > checkNum {
			curr = num
			i++
			if i > m {
				return false
			}
		}
	}
	return true
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
