package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 22:27

 * @Note:	https://leetcode.cn/problems/continuous-subarray-sum/
			523. Continuous Subarray Sum
			523. 连续的子数组和
 **/
// 暴力
//func checkSubarraySum(nums []int, k int) bool {
//	n:=len(nums)
//	preSum := make([]int,n+1)
//	for i := 0; i < n; i++ {
//		preSum[i+1]=preSum[i]+nums[i]
//	}
//	for i := 0; i < n-1; i++ {
//		for j := i+1; j < n; j++ {
//			sum:=preSum[j+1]-preSum[i]
//			if sum%k==0{
//				return true
//			}
//		}
//	}
//	return false
//}

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	sum := 0
	table := map[int]int{0: -1}
	for i := 0; i < n; i++ {
		sum += nums[i]
		// 这个 %k 的操作有点绝
		pre := sum % k
		if _, ok := table[pre]; !ok {
			table[pre] = i
		} else if i-table[pre] >= 2 {
			return true
		}
	}
	return false
}
