package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/19 16:31

 * @Note:	https://leetcode.cn/problems/count-number-of-nice-subarrays/
			1248. Count Number of Nice Subarrays
			1248. 统计「优美子数组」
 **/

func numberOfSubarrays(nums []int, k int) int {
	most := func(k int) (result int) {
		num := 0
		left := 0
		for i, v := range nums {
			if v%2 == 1 {
				num++
			}
			for num > k {
				if nums[left]%2 == 1 {
					num--
				}
				left++
			}
			//偶数
			result += i - left + 1
		}
		return
	}
	return most(k) - most(k-1)
}
