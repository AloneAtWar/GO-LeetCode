package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 19:22

 * @Note:	https://leetcode.cn/problems/sort-colors/
			75. Sort Colors
			75. 颜色分类
 **/

//func sortColors(nums []int)  {
//	// 计数排序
//	table := [3]int{}
//	for _, v := range nums {
//		table[v]++
//	}
//	index:=0
//	for i, v := range table {
//		for j := 0; j < v; j++ {
//			nums[index]=i
//			index++
//		}
//	}
//	return
//}

// 双指针
func main() {
	sortColors([]int{0, 1, 0})
}
func sortColors(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	// [0,l)  0
	// (r,len(nums)-1] 2
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			if i != l {
				swap(nums, i, l)
			}
			l++
		} else if nums[i] == 2 {
			swap(nums, r, i)
			r--
			i--
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
