package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 22:14

 * @Note:	https://leetcode.cn/problems/subarray-sum-equals-k/
			560. Subarray Sum Equals K
			560. 和为 K 的子数组
 **/
//
//func subarraySum(nums []int, k int) (result int) {
//	n:=len(nums)
//	preSum := make([]int,n+1)
//	for i := 0; i < n; i++ {
//		preSum[i+1]=preSum[i]+nums[i]
//	}
//	for i := 0; i < n+1; i++ {
//		for j := i+1; j < n+1; j++ {
//			sum:=preSum[j]-preSum[i]
//			if sum==k{
//				result++
//			}else if sum>k{
//				break
//			}
//		}
//	}
//	return
//}

// 类似two sum 的优化
func subarraySum(nums []int, k int) (result int) {
	n := len(nums)
	sum := 0
	table := map[int]int{0: 1}
	for i := 0; i < n; i++ {
		sum += nums[i]
		result += table[sum-k]
		table[sum]++
	}
	return
}
