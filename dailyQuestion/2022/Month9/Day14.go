package main

import "sort"

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/14 9:07

 * @Note:	https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/
			`1619. Mean of Array After Removing Some Elements
			1619. 删除某些元素后的数组均值
 **/

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	for _, x := range arr[n/20 : 19*n/20] {
		sum += x
	}
	return float64(sum*10) / float64(n*9)
}
