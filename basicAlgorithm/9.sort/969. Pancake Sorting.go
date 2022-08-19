package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 22:03

 * @Note:	https://leetcode.cn/problems/pancake-sorting/
			969. Pancake Sorting
			969. 煎饼排序
 **/

func pancakeSort(arr []int) (result []int) {
	curr := len(arr)
	for curr > 1 {
		index := find(arr, curr)
		if index != curr-1 {
			// 每次把最大的翻到第一个
			result = append(result, index+1, curr)
			flip(arr, index)
			flip(arr, curr-1)
		}
		curr--
	}
	return
}

func flip(arr []int, index int) {
	left := 0
	right := index
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func find(arr []int, want int) int {
	for i, v := range arr {
		if want == v {
			return i
		}
	}
	return -1
}
