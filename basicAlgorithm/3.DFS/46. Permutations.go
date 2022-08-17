package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/18 1:06

 * @Note:	https://leetcode.cn/problems/permutations/
			46. Permutations
			46. 全排列
 **/

func permute(nums []int) (result [][]int) {
	var dfs func()
	table := map[int]bool{}
	n := len(nums)
	curr := make([]int, 0, len(nums))
	dfs = func() {
		if len(curr) == n {
			result = append(result, append([]int(nil), curr...))
			return
		}
		size := len(curr)
		for i, num := range nums {
			if !table[i] {
				curr = append(curr, num)
				table[i] = true
				dfs()
				table[i] = false
				curr = curr[:size]
			}
		}
	}
	dfs()
	return
}
