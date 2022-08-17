package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 15:43

 * @Note:	https://leetcode.cn/problems/replace-elements-with-greatest-element-on-right-side/
			1299. Replace Elements with Greatest Element on Right Side
			1299. 将每个元素替换为右侧最大元素
 **/

// 白痴做法 所以扫描线 在哪里哦？？？
func replaceElements(arr []int) []int {
	n := len(arr)
	max := arr[n-1]
	result := make([]int, len(arr))
	result[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		result[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}
	return result
}
