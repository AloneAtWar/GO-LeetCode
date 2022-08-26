package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/27 2:43

 * @Note:	https://leetcode.cn/problems/find-k-closest-elements/
			658. 找到 K 个最接近的元素
			658. Find K Closest Elements
 **/

// 稳定排序 这个API 牛啊
//func findClosestElements(arr []int, k, x int) []int {
//	// 稳定排序，在绝对值相同的情况下，保证更小的数排在前面
//	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
//	arr = arr[:k]
//	sort.Ints(arr)
//	return arr
//}

// 这代码居然不超时  暴力法  自己携带
//func findClosestElements(arr []int, k int, x int)(result []int) {
//	// key 是距离 value 是 值
//	table := make(map[int][]int)
//	for _, v := range arr {
//		this := abs(v - x)
//		table[this]=append(table[this],v)
//	}
//	var list []int
//	for i, _ := range table {
//		list=append(list,i)
//	}
//	sort.Slice(list, func(i, j int) bool {
//		return list[i]<list[j]
//	})
//	index:=0
//	for k>0{
//		this:=table[list[index]]
//		addNum:=len(this)
//		if addNum>k{
//			addNum=k
//		}
//		result=append(result,this[:addNum]...)
//		k-=addNum
//		index++
//	}
//	sort.Slice(result, func(i, j int) bool {
//		return result[i]<result[j]
//	})
//	sort.SearchInts(arr, x)
//	return
//}
//
//func abs(a int)int{
//	if a>0{
//		return a
//	}
//	return -a
//}

//	0 1 0 1
//	1 0 1 0
//	0 0	1 1
//	1 1 0 0

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	right := sort.SearchInts(arr, x)
	left := right
	for ; k > 0; k-- {
		if left == 0 {
			right++
		} else if right == n || x-arr[left-1] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left:right]
}
