package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/17 21:10

 * @Note:	https://leetcode.cn/problems/subsets/
			78. Subsets
			78. 子集
 **/
func subsets(nums []int) (result [][]int) {
	var dfs func(i int)
	var curr = []int{}
	dfs = func(begin int) {
		result = append(result, append([]int(nil), curr...))
		for i := begin; i < len(nums); i++ {
			curr = append(curr, nums[i])
			dfs(i + 1)
			curr = curr[:len(curr)-1]
		}
	}
	dfs(0)
	return
}
