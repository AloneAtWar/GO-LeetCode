package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 22:35

 * @Note:	https://leetcode.cn/problems/contiguous-array/
			525. Contiguous Array
			525. 连续数组
 **/
// 暴力
//func findMaxLength(nums []int) (result int) {
//	n:=len(nums)
//	for i := 0; i < n; i++ {
//		if nums[i]==0{
//			nums[i]=-1
//		}
//	}
//	preSum := make([]int,n+1)
//	for i := 0; i < n; i++ {
//		preSum[i+1]=preSum[i]+nums[i]
//	}
//	for i := 0; i < n; i++ {
//		for j := n-1; j >=i+1  && j-i+1>result; j-- {
//			sum:=preSum[j+1]-preSum[i]
//			if sum==0{
//				result=j-i+1
//			}
//		}
//	}
//	return
//}

func findMaxLength(nums []int) (result int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	table := map[int]int{0: -1}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if j, ok := table[sum]; ok {
			if i-j > result {
				result = i - j
			}
		} else {
			table[sum] = i
		}
	}
	return
}
